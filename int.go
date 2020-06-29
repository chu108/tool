package tool

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

/**
返回范围内随机数
max: 10 最大数
min: 5 最小数
ret: 5-10 之间的数
*/
func GetRand(max, min int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max) + min
}

/**
将float64转成精确的int64
num:：数字
retain：保留位数，精度
*/
func FloatWrapInt64(num float64, retain int) int64 {
	return int64(num * math.Pow10(retain))
}

//将int64恢复成正常的float64
func Int64UnwrapFloat(num int64, retain int) float64 {
	return float64(num) / math.Pow10(retain)
}

//fload64保留两位小数
func FloatDecimal(num float64) float64 {
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", num), 64)
	return value
}

//int64转字符串
func Int64ToStr(i int64, base int) string {
	return strconv.FormatInt(i, base)
}

//int转字符串
func IntToStr(i int) string {
	return strconv.Itoa(i)
}

//float转int
func Float64ToInt(f float64) int {
	return int(f)
}

//float转int64
func Float64ToInt64(f float64) int64 {
	return int64(f)
}

/*
float转字符串
'b' (-ddddp±ddd，二进制指数)
'e' (-d.dddde±dd，十进制指数)
'E' (-d.ddddE±dd，十进制指数)
'f' (-ddd.dddd，没有指数)
'g' ('e':大指数，'f':其它情况)
'G' ('E':大指数，'f':其它情况)
*/
func FloatToStr(f float64) string {
	return strconv.FormatFloat(f, 'E', -1, 64)
}

//时间戳转日期字符串
func Int64ToDateStr(i int64) string {
	return time.Unix(i, 0).Format(DateTemp)
}

//时间戳转日期时间字符串
func Int64ToDateTimeStr(i int64) string {
	return time.Unix(i, 0).Format(TimeTemp)
}
