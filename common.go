package tool

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"os/exec"
	"strings"
)

//执行命令并返回结果
func Command(commName string, arg ...string) (string, error) {
	cmdPath, err := exec.LookPath(commName)
	if err != nil {
		panic(err)
	}
	cmd := exec.Command(cmdPath, arg...)
	output, err := cmd.CombinedOutput()
	outputStr := BytesToStr(output)
	Info(cmd.String())
	if err != nil {
		Err(cmd.String())
		return "", err
	}
	return outputStr, nil
}

//执行命令并直接输出结果
func CommandPipe(commName string, arg ...string) error {
	cmdPath, err := exec.LookPath(commName)
	if err != nil {
		panic(err)
	}
	cmd := exec.Command(cmdPath, arg...)
	// 命令的错误输出和标准输出都连接到同一个管道
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout

	if err != nil {
		return err
	}

	if err = cmd.Start(); err != nil {
		return err
	}
	// 从管道中实时获取输出并打印到终端
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		fmt.Print(string(tmp))
		if err != nil {
			break
		}
	}

	if err = cmd.Wait(); err != nil {
		return err
	}
	return nil
}

func CommandGrep(commName string, arg ...string) (string, error) {
	cmdPath, err := exec.LookPath(commName)
	if err != nil {
		return "", err
	}
	cmd := exec.Command("bash", "-c", cmdPath+" "+strings.Join(arg, " "))
	var out, stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		fmt.Println(cmd.String())
		return out.String(), errors.New(err.Error() + ":" + stderr.String())
	}
	return out.String(), nil
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
