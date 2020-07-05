package matchData

import (
	"github.com/fangker/gbdb/backend/utils/uassert"
	"math"
	"unsafe"
)

func MatchWrite1(ptr *byte, v uint) {
	*ptr = byte(v)
}

func MatchRead1(ptr *byte) uint {
	return uint(*ptr)
}
func MatchWrite2(ptr *byte, v uint) {
	uassert.True(v <= math.MaxUint16)
	ob := pointerWithByteOffset(ptr, 1)
	*ptr = byte(v >> 8)
	*ob = byte(v)
}
func MatchRead2(ptr *byte) uint {
	ob := pointerWithByteOffset(ptr, 1)
	return uint(*ptr)<<8 + uint(*ob)
}

func MatchWrite4(ptr *byte, v uint) {
	uassert.True(v <= math.MaxUint32)
	*pointerWithByteOffset(ptr, 0) = byte(v >> 24)
	*pointerWithByteOffset(ptr, 1) = byte(v >> 16)
	*pointerWithByteOffset(ptr, 2) = byte(v >> 8)
	*pointerWithByteOffset(ptr, 3) = byte(v)
}

func MatchRead4(ptr *byte) uint {
	return uint(*pointerWithByteOffset(ptr, 0))<<24 +
		uint(*pointerWithByteOffset(ptr, 1))<<16 +
		uint(*pointerWithByteOffset(ptr, 2))<<8 +
		uint(*pointerWithByteOffset(ptr, 3))
}

func MatchWrite8(ptr *byte, v uint) {
	uassert.True(v <= math.MaxUint64)
	*pointerWithByteOffset(ptr, 0) = byte(v >> 56)
	*pointerWithByteOffset(ptr, 1) = byte(v >> 48)
	*pointerWithByteOffset(ptr, 2) = byte(v >> 40)
	*pointerWithByteOffset(ptr, 3) = byte(v >> 32)
	*pointerWithByteOffset(ptr, 4) = byte(v >> 24)
	*pointerWithByteOffset(ptr, 5) = byte(v >> 16)
	*pointerWithByteOffset(ptr, 6) = byte(v >> 8)
	*pointerWithByteOffset(ptr, 7) = byte(v)
}

func MatchRead8(ptr *byte) uint {
	return uint(*pointerWithByteOffset(ptr, 0))<<56 +
		uint(*pointerWithByteOffset(ptr, 1))<<48 +
		uint(*pointerWithByteOffset(ptr, 2))<<40 +
		uint(*pointerWithByteOffset(ptr, 3))<<32 +
		uint(*pointerWithByteOffset(ptr, 4))<<24 +
		uint(*pointerWithByteOffset(ptr, 5))<<16 +
		uint(*pointerWithByteOffset(ptr, 6))<<8 +
		uint(*pointerWithByteOffset(ptr, 7))
}

func pointerWithByteOffset(ptr *byte, n uint8) *byte {
	p1 := unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + unsafe.Sizeof(*ptr)*uintptr(n))
	return (*byte)(p1)
}
