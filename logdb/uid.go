package logdb

import (
	"crypto/rand"
	"encoding/base32"
	"encoding/binary"
	"time"

	"github.com/sirupsen/logrus"
)

func genUid() string {
	timenano := uint64(time.Now().UnixNano())
	timebytes := make([]byte, 15)
	binary.BigEndian.PutUint64(timebytes[:8], timenano)
	if _, err := rand.Read(timebytes[8:]); err != nil {
		logrus.Fatalf("Failed to generate random uid: %v", err)
	}
	return base32.HexEncoding.EncodeToString(timebytes)
}
