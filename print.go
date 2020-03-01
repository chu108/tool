package tool

import "fmt"

var (
	green  = string([]byte{27, 91, 51, 50, 109})
	yellow = string([]byte{27, 91, 51, 51, 109})
	red    = string([]byte{27, 91, 51, 49, 109})
	reset  = string([]byte{27, 91, 48, 109})
)

func Err(errs ...interface{}) {
	for i := 0; i < len(errs); i++ {
		fmt.Print(red, errs[i], reset)
		//fmt.Print(fmt.Sprintf("\033[1;31m %v \033[0m", errs[i]))
	}
	fmt.Print("\n")
}

func Info(infos ...interface{}) {
	for i := 0; i < len(infos); i++ {
		fmt.Print(green, infos[i], reset)
		//fmt.Print(fmt.Sprintf("\033[1;32m %v \033[0m", infos[i]))
	}
	fmt.Print("\n")
}
