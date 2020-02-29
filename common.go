package tool

import (
	"flag"
	"os/exec"
)

//执行命令函数
func Command(commName string, param []string) string {
	cmdPath, err := exec.LookPath(commName)
	if err != nil {
		panic(err)
	}
	cmd := exec.Command(cmdPath, param...)
	output, err := cmd.CombinedOutput()
	outputStr := BytesToStr(output)
	if err != nil {
		Err(cmd.String())
		panic(outputStr)
	}
	return outputStr
}

//解析命令行字符串参数
func FlagString(name, value, usage string) string {
	val := flag.String(name, value, usage)
	flag.Parse()
	return *val
}

//解析命令行int参数
func FlagInt(name string, value int, usage string) int {
	val := flag.Int(name, value, usage)
	flag.Parse()
	return *val
}

//解析命令行int64参数
func FlagInt64(name string, value int64, usage string) int64 {
	val := flag.Int64(name, value, usage)
	flag.Parse()
	return *val
}

//解析命令行int64参数
func FlagFloat64(name string, value int64, usage string) float64 {
	val := flag.Float64(name, 0, usage)
	flag.Parse()
	return *val
}

//解析命令行bool参数
func FlagBool(name string, value bool, usage string) bool {
	val := flag.Bool(name, value, usage)
	flag.Parse()
	return *val
}
