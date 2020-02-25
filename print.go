package tool

import "fmt"

func Err(errs ...interface{}) {
	for i := 0; i < len(errs); i++ {
		fmt.Print(fmt.Sprintf("\033[1;32m %v \033[0m", errs[i]))
	}
	fmt.Print("\n")
}

func Info(infos ...interface{}) {
	for i := 0; i < len(infos); i++ {
		fmt.Print(fmt.Sprintf("\033[1;32m %v \033[0m", infos[i]))
	}
	fmt.Print("\n")
}
