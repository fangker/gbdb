package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"sync/atomic"
	"time"
)

type UUID []byte

var machineId = readMachineId()
var incrementCount uint32 = 0

func readMachineId() []byte {
	var mid [3]byte
	hostname, nohosterr := os.Hostname()
	if nohosterr != nil {
		_, randerr := io.ReadFull(rand.Reader, mid[:])
		if randerr != nil {
			panic(fmt.Errorf("can't get machineId"))
		}
	}
	hw := md5.New()
	hw.Write([]byte(hostname))
	copy(mid[:], hw.Sum(nil))
	return mid[:]
}

//    UUID  |4 timestamp| 3 machineId |  2 PID | 3 Increment num|
func NewUUID() UUID {
	var uuid [12]byte
	binary.BigEndian.PutUint32(uuid[:], uint32(time.Now().Unix()))
	copy(uuid[3:6], machineId[:])
	pid := os.Getpid()
	uuid[7], uuid[8] = byte(pid>>8), byte(pid)
	i := atomic.AddUint32(&incrementCount, 1)
	uuid[9] = byte(i >> 16)
	uuid[10] = byte(i >> 8)
	uuid[11] = byte(i)
	return UUID(hex.EncodeToString(uuid[:]))
}
