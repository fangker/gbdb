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
	return &mtrLog{buf:bytes.NewBuffer([]byte{})}
}

func MLogClose(mtr *Mtr,ml *mtrLog) {
	mtr.nLogRecs++
	_, _ = mtr.log.ReadFrom(ml.buf)
}

// mLog open 载入用于写入
func (ml *mtrLog) MLogWriteUint(ptr *byte, v uint32, logType MLOG_TYPE) {
	if logType == MLOG_TYPE_BYRE_1 {
		ml.buf.WriteByte(byte(MLOG_TYPE_BYRE_1))
		binary.Write(ml.buf, binary.BigEndian, uint8(v))
	}
	if logType == MLOG_TYPE_BYRE_2 {
		ml.buf.WriteByte(byte(MLOG_TYPE_BYRE_2))
		binary.Write(ml.buf, binary.BigEndian, uint16(v))
	}
	if logType == MLOG_TYPE_BYRE_4 {
		ml.buf.WriteByte(byte(MLOG_TYPE_BYRE_4))
		binary.Write(ml.buf, binary.BigEndian, uint32(v))
	}
}

func (ml *mtrLog) MLogWriteDUint(ptr *byte, v uint64) {
	ml.buf.WriteByte(byte(MLOG_TYPE_BYRE_8))
	binary.Write(ml.buf, binary.BigEndian, uint64(v))
}

func (ml *mtrLog) MLogWriteString(ptr *byte, bs []byte) {
	ml.buf.WriteByte(byte(MLOG_TYPE_BYRE_STRING))
	binary.Write(ml.buf, binary.BigEndian, uint16(len(bs)))
	binary.Write(ml.buf, binary.BigEndian, bs)
}

func MLogInitialRecord(ptr *byte, ml *mtrLog) *mtrLog {
	bp := cachehelper.PosInBlockAlign(ptr)
	offset := cachehelper.OffsetInBlockAlign(ptr)
	spaceId, pageOffset := bp.GetPos()
	binary.Write(ml.buf, binary.BigEndian, uint16(spaceId))
	binary.Write(ml.buf, binary.BigEndian, uint16(pageOffset))
	binary.Write(ml.buf, binary.BigEndian, uint16(offset))
	return ml
}
