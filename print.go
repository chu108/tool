package tool

import "fmt"

var (
	green  = string([]byte{27, 91, 51, 50, 109})
	yellow = string([]byte{27, 91, 51, 51, 109})
	red    = string([]byte{27, 91, 51, 49, 109})
	reset  = string([]byte{27, 91, 48, 109})
)

func Err(errs ...interface{}) {
	fmt.Printf("%s %v %s \n", red, errs, reset)
}

func Info(infos ...interface{}) {
	fmt.Printf("%s %v %s \n", green, infos, reset)
}
