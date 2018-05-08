package cache

import "os"

type Wrapper struct {
	TableID uint32
	File    *os.File
}
