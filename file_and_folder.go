package tool

import (
	"bufio"
	"io"
	"os"
)

//如果文件夹不存在，则递归创建文件夹
func CreateFileByNot(filePath string) error {
	if !IsExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		return err
	}
	return nil
}

// 判断所给路径文件/文件夹是否存在(返回true是存在)
func IsExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

/**
一次性加载文件，按行读取
*/
func ReadFileForReader(fielPath string, callBak func(row string) bool) {
	file, err := os.OpenFile(fielPath, os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := bufio.NewReader(file)

	for {
		b, err := buf.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		rowStr := BytesToStr(b)
		if rowStr != "" && !callBak(rowStr) {
			break
		}
	}
}

/**
逐行读取文件
*/
func ReadFileForScanner(fielPath string, callBak func(row string) bool) {
	file, err := os.OpenFile(fielPath, os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		rowStr := BytesToStr(scanner.Bytes())
		if rowStr != "" && !callBak(rowStr) {
			break
		}
	}
}
