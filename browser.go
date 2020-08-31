package tool

import (
	"os/exec"
	"runtime"
)

//打开网页
func OpenBrowser(url string) (err error) {
	switch runtime.GOOS {
	case "windows": //windows
		err = exec.Command(`cmd`, `/c`, `start`, url).Start()
	case "darwin": //Linux
		err = exec.Command(`xdg-open`, url).Start()
	default: //Mac
		err = exec.Command(`open`, url).Start()
	}
	return
}
