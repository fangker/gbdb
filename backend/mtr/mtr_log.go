package mtr

import (
	"bytes"
	"encoding/binary"
	"github.com/fangker/gbdb/backend/cache/cachehelper"
)

const (
	MLOG_TYPE_BYRE_1      MLOG_TYPE = 0
	MLOG_TYPE_BYRE_2      MLOG_TYPE = 1
	MLOG_TYPE_BYRE_4      MLOG_TYPE = 2
	MLOG_TYPE_BYRE_8      MLOG_TYPE = 3
	MLOG_TYPE_BYRE_STRING MLOG_TYPE = 4
)

type MLOG_TYPE byte

// Mtr Logs
type mtrLog struct {
	buf *bytes.Buffer
}

func MLogOpen() *mtrLog {
	return &mtrLog{}
}

func (mtrlog *mtrLog) MLogClose(mtr *Mtr) {
	mtr.nLogRecs++
	mtr.log.ReadFrom(mtrlog.buf);
}

// mlog open 载入用于写入
func (mtrlog *mtrLog) MLogWriteUint(ptr *byte, v uint32, logType MLOG_TYPE) {
	mtrlog.buf.WriteByte(byte(logType))
	if logType == MLOG_TYPE_BYRE_1 {
		binary.Write(mtrlog.buf, binary.BigEndian, uint8(v))
	}
	if logType == MLOG_TYPE_BYRE_2 {
		binary.Write(mtrlog.buf, binary.BigEndian, uint16(v))
	}
	if logType == MLOG_TYPE_BYRE_4 {
		binary.Write(mtrlog.buf, binary.BigEndian, uint32(v))
	}
}

func (mtrlog *mtrLog) MLogWriteDUint(ptr *byte, v uint64, logType MLOG_TYPE) {
	mtrlog.buf.WriteByte(byte(logType))
	if logType == MLOG_TYPE_BYRE_8 {
		binary.Write(mtrlog.buf, binary.BigEndian, uint64(v))
	}
}
func MLogInitialRecord(ptr *byte) {
	cachehelper.BlockPageAlign(ptr)
}
