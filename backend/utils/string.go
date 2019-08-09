package utils

import "unsafe"

// github_henrylee2cn_util
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func StringToBytes(s string) []byte {
	// ptr length
	sp := *(*[2]uintptr)(unsafe.Pointer(&s))
	// ptr len cap
	bp := [3]uintptr{sp[0], sp[1], sp[1]}
	return *(*[]byte)(unsafe.Pointer(&bp))
}

