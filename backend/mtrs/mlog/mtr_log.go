package mlog

import (
	"bytes"
	"encoding/binary"
	"github.com/fangker/gbdb/backend/cache/cachehelper"
	. "github.com/fangker/gbdb/backend/dm/matchData"
	. "github.com/fangker/gbdb/backend/mtrs/mtr"
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

func Open() *mtrLog {
	return &mtrLog{buf: bytes.NewBuffer([]byte{})}
}

func Close(mtr *Mtr, ml *mtrLog) {
	MergeMLog(mtr, ml.buf)
}

// mLog open 载入用于写入
func WriteUint(ptr *byte, v uint, logType MLOG_TYPE) {
	ml := Open()
	var into interface{}
	if logType == MLOG_TYPE_BYRE_1 {
		MatchWrite1(ptr, v)
		into = uint8(v)
	}
	if logType == MLOG_TYPE_BYRE_2 {
		MatchWrite2(ptr, v)
		into = uint16(v)
	}
	if logType == MLOG_TYPE_BYRE_4 {
		MatchWrite4(ptr, v)
		into = uint32(v)
	}
	if logType == MLOG_TYPE_BYRE_8 {
		MatchWrite8(ptr, v)
		into = uint64(v)
	}
	ml.buf.WriteByte(byte(logType))
	binary.Write(ml.buf, binary.BigEndian, into)
}

//func ParseUint(ptr *byte, uint offset, *byte page) {
//
//}

func (ml *mtrLog) MLogWriteString(ptr *byte, bs []byte) {
	ml.buf.WriteByte(byte(MLOG_TYPE_BYRE_STRING))
	binary.Write(ml.buf, binary.BigEndian, uint16(len(bs)))
	binary.Write(ml.buf, binary.BigEndian, bs)
}

func InitialRecord(ptr *byte, ml *mtrLog) *mtrLog {
	bp := cachehelper.PosInBlockAlign(ptr)
	offset := cachehelper.OffsetInBlockAlign(ptr)
	spaceId, pageOffset := bp.GetPos()
	binary.Write(ml.buf, binary.BigEndian, uint16(spaceId))
	binary.Write(ml.buf, binary.BigEndian, uint16(pageOffset))
	binary.Write(ml.buf, binary.BigEndian, uint16(offset))
	return ml
}
