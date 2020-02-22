package tool

import "math/rand"

/**
返回范围内随机数
max: 10 最大数
min: 5 最小数
ret: 5-10 之间的数
*/
func GetRand(max, min int) int {
	return rand.Intn(max) + min
}
