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

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	pb "matsu.dev/matsuri-contest-log/proto"
)

const (
	MinSupportedDatabaseVersion = 1
	DatabaseVersion             = 1
)

type QSO struct {
	Uid           string
	LocalCallsign string
	DXCallsign    string
	Time          int64
	Freq          int64
	Mode          string
	ExchangeSent  map[string]string
	ExchangeRcvd  map[string]string
	Type          pb.QSOType
	Operator      string
}

type snapshotEvaluator struct {
	activeQso  map[string]*QSO
	stagingQso map[string]*QSO
}

type BinlogServer struct {
	nextSerialNumber uint64
	mutex            sync.RWMutex

	currentContest *pb.ActiveContest
	db             *sql.DB
	evaluator      *snapshotEvaluator
}

func (s *BinlogServer) rlock() {
	s.mutex.RLock()
}

func (s *BinlogServer) runlock() {
	s.mutex.RUnlock()
}

func (s *BinlogServer) lock() {
	s.mutex.Lock()
}

func (s *BinlogServer) unlock() {
	s.mutex.Unlock()
}

func FromProtobufQso(v *pb.QSO) QSO {
	return QSO{
		Uid:           v.Uid,
		LocalCallsign: v.LocalCallsign,
		DXCallsign:    v.DxCallsign,
		Time:          v.Time,
		Freq:          v.Freq,
		Mode:          v.Mode,
		ExchangeSent:  v.ExchangeSent,
		ExchangeRcvd:  v.ExchangeRcvd,
		Type:          v.Type,
		Operator:      v.Operator,
	}
}

func (q *QSO) ToProtobufQso() *pb.QSO {
	return &pb.QSO{
		Uid:           q.Uid,
		LocalCallsign: q.LocalCallsign,
		DxCallsign:    q.DXCallsign,
		Time:          q.Time,
		Freq:          q.Freq,
		Mode:          q.Mode,
		ExchangeSent:  q.ExchangeSent,
		ExchangeRcvd:  q.ExchangeRcvd,
		Type:          q.Type,
		Operator:      q.Operator,
	}
}

func (s *BinlogServer) Push(ctx context.Context, messages *pb.BinlogMessageSet) (*pb.StandardResponse, error) {
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

func (s *BinlogServer) Retrieve(ctx context.Context, req *pb.RetrieveBinlogRequest) (*pb.BinlogMessageSet, error) {
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

func (s *BinlogServer) RetrieveSnapshot(ctx context.Context, req *pb.RetrieveBinlogRequest) (*pb.SnapshotMessage, error) {
	ret := &pb.SnapshotMessage{}
	ret.Qso = make([]*pb.QSO, 0)

	if req.SerialEnd == 0x7fff_ffff_ffff_ffff {
		for _, qso := range s.evaluator.activeQso {
			ret.Qso = append(ret.Qso, qso.ToProtobufQso())
		}
		ret.Serial = s.nextSerialNumber
		return ret, nil
	}

	res := s.db.QueryRow("SELECT serial, content FROM snapshot WHERE serial >= ? AND serial < ? ORDER BY serial DESC LIMIT 1", req.GetSerialStart(), req.GetSerialEnd())
	if res == nil {
		return &pb.SnapshotMessage{}, nil
	}

	var (
		data   []byte
		serial uint64
	)
	if err := res.Scan(&serial, &data); err != nil {
		return &pb.SnapshotMessage{Qso: []*pb.QSO{}, Serial: 0}, nil
	}
	proto.Unmarshal(data, ret)
	ret.Serial = serial
	return ret, nil
}

func (s *BinlogServer) BuildSnapshot() {
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
		snapshot.Qso = append(snapshot.Qso, v.ToProtobufQso())
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
		qso := FromProtobufQso(val.Qso)
		if qso.Uid == "" {
			return errors.New("missing uid in QSO")
		}
		e.stagingQso[qso.Uid] = &qso
	}
	return nil
}

func NewContest(filename string, manifest *pb.ActiveContest) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	logrus.Infof("opened database: %v", filename)

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

	protoString, _ := protojson.Marshal(manifest)

	values := map[string]string{
		"contest":              string(protoString),
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

func LoadContest(filename string) (*sql.DB, *pb.ActiveContest, error) {
	manifest := &pb.ActiveContest{}
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return nil, nil, err
	}

	close_db := db
	defer func() {
		if close_db != nil {
			close_db.Close()
		}
	}()

	rows, err := db.Query(`SELECT key, value FROM kvstorage`)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to initialize the database: %v", err)
	}

	for rows.Next() {
		var (
			k string
			v string
		)

		rows.Scan(&k, &v)

		switch k {
		case "contest":
			err := protojson.Unmarshal([]byte(v), manifest)
			if err != nil {
				return nil, nil, fmt.Errorf("contest is not valid: %v", err)
			}
		case "database_version":
			if databaseVersion, err := strconv.Atoi(v); databaseVersion > DatabaseVersion {
				logrus.Warnf("database contains data of newer schema version (version %d), some features might be unsupported by this version (version %d), please consider upgrade your MCL", databaseVersion, DatabaseVersion)
			} else if err != nil {
				return nil, nil, fmt.Errorf("cannot load contest: invalid version: %v", err)
			}
		case "min_database_version":
			var min_version int
			min_version, err = strconv.Atoi(v)

			if min_version > DatabaseVersion {
				return nil, nil, fmt.Errorf("database is not supported by this version of MCL (schema version %d, the database requires minimal version of %d), please upgrade your MCL", DatabaseVersion, min_version)
			} else if err != nil {
				return nil, nil, fmt.Errorf("cannot load contest: invalid version: %v", err)
			}
		}
	}

	close_db = nil

	return db, manifest, nil
}

func (s *BinlogServer) RecoverFromBinlog() error {
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
	Database string
}

func NewBinlogServer(conf *BinlogServerConfig) (*BinlogServer, error) {
	mclServ := &BinlogServer{
		evaluator:      NewSnapshotEvaluator(),
		currentContest: &pb.ActiveContest{},
	}

	var err error

	mclServ.db, mclServ.currentContest, err = LoadContest(conf.Database)
	if err != nil {
		return nil, err
	}

	startTime := time.Now()
	if err := mclServ.RecoverFromBinlog(); err != nil {
		logrus.Panicf("failed to load database (probably broken): %v", err)
	}
	logrus.Infof("loading QSO binlog costs %dms", time.Since(startTime).Milliseconds())

	mclServ.BuildSnapshot()

	return mclServ, nil
}

func (s *BinlogServer) Shutdown() {
	s.db.Close()
}
