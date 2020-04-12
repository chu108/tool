package tool

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

/**
返回范围内随机数
max: 10 最大数
min: 5 最小数
ret: 5-10 之间的数
*/
func GetRand(max, min int) int {
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
