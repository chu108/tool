package tool

import "os/exec"

//执行命令函数
func command(commName string, param []string) string {
	cmdPath, _ := exec.LookPath(commName)
	cmd := exec.Command(cmdPath, param...)
	output, err := cmd.CombinedOutput()
	outputStr := BytesToStr(output)
	if err != nil {
		Err(cmd.String())
		panic(outputStr)
	}
	return outputStr
}
