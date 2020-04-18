package tool

import (
	"bufio"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
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

/**
获取当前文件夹中的所有文件
*/
func ReadDirFiles(dir string) ([]string, error) {
	fileList := make([]string, 0, 30)
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && info.Name()[:1] != "." {
			fileList = append(fileList, dir+info.Name())
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return fileList, nil
}

/**
递归获取当前文件夹中的所有文件与子文件
*/
func ReadDirFilesAll() {

}

/*
获取文件大小
*/
func GetFileSize(file string) int64 {
	fileInfo, err := os.Stat(file)
	if err != nil {
		return 0
	}
	return fileInfo.Size()
}

func ConvertByte2String(byte []byte, charset Charset) string {

	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}

/**
一次性读取文件所有内容
*/
func ReadFile(fielPath string) ([]byte, error) {
	file, err := os.OpenFile(fielPath, os.O_RDWR, 0777)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return ioutil.ReadAll(file)
}
