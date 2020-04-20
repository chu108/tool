package tool

import (
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
	"unsafe"
)

const DateTemp = "2006-01-02"
const TimeTemp = "2006-01-02 15:04:05"

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

//字符串拆分，按字数
func StrSplitByNum(txt string, length int) []string {
	txtRune := []rune(txt)
	txtLen := len(txtRune)
	retTxt := make([]string, 0, 5)
	if txtLen > length {
		for i := 0; i < txtLen; i += length {
			retTxt = append(retTxt, string(txtRune[i:i+length]))
		}
	} else {
		retTxt = append(retTxt, txt)
	}
	return retTxt
}

//字符串转int
func StrToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

//字符串转int64
func StrToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

//字符串转float
func StrToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

//int字符串转日期字符串
func StrToDateStr(s string) (string, error) {
	i, err := StrToInt64(s)
	if err != nil {
		return "", err
	}
	return time.Unix(i, 0).Format(DateTemp), nil
}

//int字符串转时间字符串
func StrToDateTimeStr(s string) (string, error) {
	i, err := StrToInt64(s)
	if err != nil {
		return "", err
	}
	return time.Unix(i, 0).Format(TimeTemp), nil
}

//日期字符串转时间对象
func StrToTime(s string) (time.Time, error) {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return time.Time{}, err
	}
	return time.ParseInLocation(TimeTemp, s, loc)
}
