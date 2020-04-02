package tool

import "fmt"

/**
格式：\033[显示方式;前景色;背景色m

说明：
前景色            背景色           颜色
---------------------------------------
30                40              黑色
31                41              红色
32                42              绿色
33                43              黃色
34                44              蓝色
35                45              紫红色
36                46              青蓝色
37                47              白色
显示方式           意义
-------------------------
0                终端默认设置
1                高亮显示
4                使用下划线
5                闪烁
7                反白显示
8                不可见

例子：
\033[1;31;40m
\033[0m
*/

const (
	red = uint8(iota + 31)
	green
	yellow
	blue
	magenta
	cyan
	white
)

const printCode = "\x1b[%dm%s\x1b[0m \n"

func Err(errs ...interface{}) {
	fmt.Printf(printCode, red, errs)
}

func Info(infos ...interface{}) {
	fmt.Printf(printCode, green, infos)
}
