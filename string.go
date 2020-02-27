package tool

import (
	"regexp"
	"strings"
	"unicode/utf8"
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

//截取字符串
func Substr(str string, length int) string {
	runeStr := []rune(str)
	runeLen := len(runeStr)
	if runeLen > length {
		str = string(runeStr[runeLen-length:])
	}
	return str
}

//高效截取字符串
func SubStrDecode(str string, length int) string {
	var size, n int
	for i := 0; i < length && n < len(str); i++ {
		_, size = utf8.DecodeRuneInString(str[n:])
		n += size
	}

	return str[:n]
}

//字符串替换
func ReplaceStrEmpty(str string, rep ...string) string {
	for i := 0; i < len(rep); i++ {
		str = strings.Replace(str, rep[i], "", -1)
	}
	return str
}

//字符串正则替换
func ReplaceRegexpStrEmpty(str string, math ...string) string {
	for i := 0; i < len(math); i++ {
		rep := regexp.MustCompile(math[i])
		str = strings.Replace(str, rep.FindString(str), "", -1)
	}
	return str
}
