package tool

import (
	"strings"
	"unsafe"
)

//字符串转byte
func StrToBytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

//byte转字符串
func BytesToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

//拼接字符串
func JoinStr(strs ...string) string {
	var buider strings.Builder
	for i := 0; i < len(strs); i++ {
		buider.WriteString(strs[i])
	}
	return buider.String()
}
