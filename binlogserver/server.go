package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	pb "matsu.dev/matsuri-contest-log/proto"
)

const (
	MinSupportedDatabaseVersion = 1
	DatabaseVersion             = 1
)

type ContextManifest struct {
	Uid                string    `json:"uid"`
	Name               string    `json:"name"`
	Filename           string    `json:"filename"`
	Category           string    `json:"catrgory"`
	BeginTimestamp     time.Time `json:"begin_timestamp"`
	EndTimestamp       time.Time `json:"end_timestamp"`
	ContestDisplayName string    `json:"contest_display_name"`
	DatabaseVersion    int       `json:"database_version"`
}

type QSO struct {
	Uid           string
	LocalCallsign string
	DXCallsign    string
	Time          int64
	Freq          int64
	Mode          pb.Mode
	IsSatellite   bool
	ExchSent      []string
	ExchRcvd      []string
	Type          pb.QSOType
	Operator      string
}

type snapshotEvaluator struct {
	activeQso  map[string]*QSO
	stagingQso map[string]*QSO
}

type MclServer struct {
	pb.UnimplementedBinlogServer

	nextSerialNumber uint64
	mutex            sync.RWMutex

	currentContest ContextManifest
	db             *sql.DB
	evaluator      *snapshotEvaluator
}

func (s *MclServer) rlock() {
	s.mutex.RLock()
}

func (s *MclServer) runlock() {
	s.mutex.RUnlock()
}

func (s *MclServer) lock() {
	s.mutex.Lock()
}

func (s *MclServer) unlock() {
	s.mutex.Unlock()
}

func (s *MclServer) Push(ctx context.Context, messages *pb.BinlogMessageSet) (*pb.StandardResponse, error) {
	s.lock()
	defer s.unlock()

	if messages == nil {
		return &pb.StandardResponse{
			ResultCode:   pb.ResultCode_success,
			ErrorMessage: "null message",
		}, nil
	}

	expectedSerial := s.nextSerialNumber
	stagingSnapshot := s.evaluator.Copy()

	for idx, message := range messages.GetMessages() {
		if message == nil {
			continue
		}

		if message.GetSerial() != expectedSerial {
			return &pb.StandardResponse{
				ResultCode:   1,
				ErrorMessage: fmt.Sprintf("message %d has unexpected serial number, expected %d", idx, expectedSerial),
			}, nil
		}

		expectedSerial++
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to prepare database transaction: %s", err.Error())
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, "INSERT INTO binlog (serial, content) VALUES (?, ?)")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to prepare database transaction: %s", err.Error())
	}

	for idx, message := range messages.GetMessages() {
		if message == nil {
			continue
		}

		data, err := proto.Marshal(message)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to save binlog %d into database: %v", idx, err.Error())
		}

		if err := stagingSnapshot.Eval(message); err != nil {
			return &pb.StandardResponse{
				ResultCode:   pb.ResultCode_success,
				ErrorMessage: fmt.Sprintf("failed to evaluate binlog %d: %v", idx, err),
			}, nil
		}

		if _, err := stmt.ExecContext(ctx, message.GetSerial(), data); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to save binlog %d into database: %v", idx, err.Error())
		}
	}

	tx.Commit()
	s.nextSerialNumber = expectedSerial
	s.evaluator = stagingSnapshot

	logrus.Infof("database now have %d active QSOs", int(len(stagingSnapshot.activeQso)))

	return &pb.StandardResponse{
		ResultCode:   pb.ResultCode_success,
		ErrorMessage: fmt.Sprintf("success, next log serial: %d", expectedSerial),
	}, nil
}

func (s *MclServer) Retrieve(ctx context.Context, req *pb.RetrieveBinlogRequest) (*pb.BinlogMessageSet, error) {
	ret := &pb.BinlogMessageSet{}
	ret.Messages = make([]*pb.BinlogMessage, 0)

	res, err := s.db.Query("SELECT content FROM binlog WHERE serial >= ? AND serial < ?", req.GetSerialStart(), req.GetSerialEnd())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to retrieve data: %v", err)
	}

	for res.Next() {
		var data []byte
		if err = res.Scan(&data); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to retrieve data: %v", err)
		}
		parsed := &pb.BinlogMessage{}
		proto.Unmarshal(data, parsed)
		ret.Messages = append(ret.Messages, parsed)
	}
	return ret, nil
}

func (s *MclServer) RetrieveSnapshot(ctx context.Context, req *pb.RetrieveBinlogRequest) (*pb.SnapshotMessage, error) {
	ret := &pb.SnapshotMessage{}
	ret.Qso = make([]*pb.QSO, 0)

	res := s.db.QueryRow("SELECT serial, content FROM snapshot WHERE serial >= ? AND serial < ? ORDER BY serial DESC LIMIT 1", req.GetSerialStart(), req.GetSerialEnd())
	if res == nil {
		return &pb.SnapshotMessage{}, nil
	}

	var (
		data   []byte
		serial uint64
	)
	if err := res.Scan(&serial, &data); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to retrieve data: %v", err)
	}
	proto.Unmarshal(data, ret)
	ret.Serial = serial
	return ret, nil
}

func (s *MclServer) BuildSnapshot() {
	var (
		data []byte
		err  error
	)

	snapshot := &pb.SnapshotMessage{}
	res := s.db.QueryRow("SELECT serial FROM snapshot ORDER BY serial DESC LIMIT 1")

	var (
		maxSerial int
	)

	if err := res.Scan(&maxSerial); err != nil && err != sql.ErrNoRows {
		logrus.Errorf("failed to get snapshot: %v", err)
		return
	}

	if maxSerial == int(s.nextSerialNumber) {
		logrus.Infof("existing snapshot %d is the latest", maxSerial)
		return
	}

	for _, v := range s.evaluator.activeQso {
		snapshot.Qso = append(snapshot.Qso, &pb.QSO{
			Uid:           v.Uid,
			LocalCallsign: v.LocalCallsign,
			DxCallsign:    v.DXCallsign,
			Time:          v.Time,
			Freq:          v.Freq,
			Mode:          v.Mode,
			IsSatellite:   v.IsSatellite,
			ExchSent:      v.ExchSent,
			ExchRcvd:      v.ExchRcvd,
			Type:          v.Type,
			Operator:      v.Operator,
		})
	}

	if data, err = proto.Marshal(snapshot); err != nil {
		logrus.Errorf("failed to marshal snapshot: %v", err)
		return
	}

	if _, err := s.db.Exec("INSERT INTO snapshot (serial, content) VALUES (?, ?)", s.nextSerialNumber, data); err != nil {
		logrus.Errorf("failed to save snapshot: %v", err)
	}
}

func copyQsoMap(mp map[string]*QSO) map[string]*QSO {
	ret := make(map[string]*QSO, len(mp))
	for k, v := range mp {
		ret[k] = v
	}
	return ret
}

func NewSnapshotEvaluator() *snapshotEvaluator {
	return &snapshotEvaluator{
		activeQso:  make(map[string]*QSO, 0),
		stagingQso: make(map[string]*QSO, 0),
	}
}

func (e *snapshotEvaluator) Copy() *snapshotEvaluator {
	return &snapshotEvaluator{
		activeQso:  copyQsoMap(e.activeQso),
		stagingQso: copyQsoMap(e.stagingQso),
	}
}

func (e *snapshotEvaluator) Eval(msg *pb.BinlogMessage) error {
	body := msg.GetMessage()
	switch val := body.(type) {
	case *pb.BinlogMessage_QsoOp:
		switch val.QsoOp.GetType() {
		case pb.QSOOperationType_add_or_update:
			if qso, ok := e.stagingQso[val.QsoOp.GetUid()]; !ok {
				return fmt.Errorf("QSO %s not found", val.QsoOp.GetUid())
			} else {
				e.activeQso[val.QsoOp.GetUid()] = qso
			}
		case pb.QSOOperationType_delete:
			if _, ok := e.activeQso[val.QsoOp.GetUid()]; !ok {
				return fmt.Errorf("QSO %s not found", val.QsoOp.GetUid())
			} else {
				delete(e.activeQso, val.QsoOp.GetUid())
			}
		default:
			return errors.New("unknown QSO operation")
		}
	case *pb.BinlogMessage_Qso:
		qso := &QSO{
			Uid:           val.Qso.GetUid(),
			LocalCallsign: val.Qso.GetLocalCallsign(),
			DXCallsign:    val.Qso.GetDxCallsign(),
			Time:          val.Qso.GetTime(),
			Freq:          val.Qso.GetFreq(),
			Mode:          val.Qso.GetMode(),
			IsSatellite:   val.Qso.GetIsSatellite(),
			ExchSent:      val.Qso.GetExchSent(),
			ExchRcvd:      val.Qso.GetExchRcvd(),
			Type:          val.Qso.GetType(),
			Operator:      val.Qso.GetOperator(),
		}
		if qso.Uid == "" {
			return errors.New("missing uid in QSO")
		}
		e.stagingQso[qso.Uid] = qso
	}
	return nil
}

func NewContest(manifest ContextManifest) {
	db, err := sql.Open("sqlite3", manifest.Filename)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	logrus.Infof("opened database: %v", manifest.Filename)

	sqlStmt := `
	CREATE TABLE binlog (serial INTEGER NOT NULL PRIMARY KEY, content BLOB);
	DELETE FROM binlog;
	CREATE TABLE snapshot (serial INTEGER NOT NULL PRIMARY KEY, content BLOB);
	DELETE FROM snapshot;
	CREATE TABLE kvstorage (key TEXT PRIMARY KEY, value TEXT);
	DELETE FROM kvstorage;
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		panic(fmt.Errorf("failed to initialize the database: %v", err))
	}

	stmt, err := db.Prepare(`INSERT INTO kvstorage (key, value) VALUES (?, ?)`)
	if err != nil {
		panic(fmt.Errorf("failed to initialize the database: %v", err))
	}

	values := map[string]string{
		"uid":                  manifest.Uid,
		"name":                 manifest.Name,
		"category":             manifest.Category,
		"begin":                manifest.BeginTimestamp.Format(time.RFC3339),
		"end":                  manifest.EndTimestamp.Format(time.RFC3339),
		"contest_display_name": manifest.ContestDisplayName,
		"database_version":     strconv.Itoa(DatabaseVersion),
		"min_database_version": strconv.Itoa(MinSupportedDatabaseVersion),
	}

	for k, v := range values {
		_, err := stmt.Exec(k, v)
		if err != nil {
			panic(fmt.Errorf("failed to initialize the database: %v", err))
		}
	}
}

func (manifest *ContextManifest) LoadContest() *sql.DB {
	db, err := sql.Open("sqlite3", manifest.Filename)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(`SELECT key, value FROM kvstorage`)
	if err != nil {
		panic(fmt.Errorf("failed to initialize the database: %v", err))
	}

	values := map[string]*string{
		"uid":                  &manifest.Uid,
		"name":                 &manifest.Name,
		"category":             &manifest.Category,
		"contest_display_name": &manifest.ContestDisplayName,
	}

	for rows.Next() {
		var (
			k string
			v string
		)

		rows.Scan(&k, &v)

		if vptr, ok := values[k]; ok {
			*vptr = v
		} else {
			switch k {
			case "begin":
				manifest.BeginTimestamp, err = time.Parse(time.RFC3339, v)
			case "end":
				manifest.EndTimestamp, err = time.Parse(time.RFC3339, v)
			case "database_version":
				manifest.DatabaseVersion, err = strconv.Atoi(v)

				if manifest.DatabaseVersion > DatabaseVersion {
					logrus.Warnf("database contains data of newer schema version (version %d), some features might be unsupported by this version (version %d), please consider upgrade your MCL", manifest.DatabaseVersion, DatabaseVersion)
				}
			case "min_database_version":
				var min_version int
				min_version, err = strconv.Atoi(v)

				if min_version > DatabaseVersion {
					logrus.Panicf("database is not supported by this version of MCL (schema version %d, the database requires minimal version of %d), please upgrade your MCL", DatabaseVersion, min_version)
				}
			}
		}
		if err != nil {
			logrus.Panicf("failed to load the database: %v", err)
		}
	}

	return db
}

func (s *MclServer) RecoverFromBinlog() error {
	res, err := s.db.Query(`SELECT serial, content FROM binlog ORDER BY serial ASC`)
	if err != nil {
		return err
	}

	rows := 0

	for res.Next() {
		rows++
		var (
			serial uint64
			data   []byte
		)

		if err := res.Scan(&serial, &data); err != nil {
			return err
		}

		var msg pb.BinlogMessage
		proto.Unmarshal(data, &msg)

		s.nextSerialNumber = serial + 1

		if err := s.evaluator.Eval(&msg); err != nil {
			return err
		}
	}

	logrus.Infof("loaded %d items from database", rows)
	logrus.Infof("database now have %d active QSOs", len(s.evaluator.activeQso))
	return nil
}

type BinlogServerConfig struct {
	GrpcServer *grpc.Server
	Database   string
}

func NewServer(conf *BinlogServerConfig) *MclServer {
	mclServ := &MclServer{
		evaluator: NewSnapshotEvaluator(),
		currentContest: ContextManifest{
			Filename: conf.Database,
		},
	}

	mclServ.db = mclServ.currentContest.LoadContest()

	{
		startTime := time.Now()
		if err := mclServ.RecoverFromBinlog(); err != nil {
			logrus.Panicf("failed to load database (probably broken): %v", err)
		}
		logrus.Infof("loading QSO binlog costs %dms", time.Now().Sub(startTime).Milliseconds())
	}

	mclServ.BuildSnapshot()
	pb.RegisterBinlogServer(conf.GrpcServer, mclServ)

	return mclServ
}

func (s *MclServer) Init(conf *BinlogServerConfig) {
	s.BuildSnapshot()
}
