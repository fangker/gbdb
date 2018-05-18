package redo

import (
	"os"

	"github.com/fangker/gbdb/backend/dm/constants/cType"
)

type (
	Lsn cType.Lsn
)

type logGroup struct {
	nFiles    int
	file      []*os.File
	fileSize  Lsn
	lsnOffset Lsn
}
