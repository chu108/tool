package tool

import (
	"bytes"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

/**
下载文件
url：下载地址
savePath：保存路径，不包含文件名
saveName：保存文件名，如果为空，则用原名
*/
func DownloadFile(url, savePath, saveName string) error {
	if url == "" || savePath == "" {
		return errors.New("下载url或保存地址错误")
	}
	//获取文件类型
	fileExt := path.Ext(url)
	//获取文件名
	fileName := path.Base(url)
	if saveName == "" {
		saveName = fileName
	}

	if savePath[len(savePath)-1:] != "/" {
		savePath += "/"
	}

	body, err := GetByteForUrl(url)
	if err != nil {
		return err
	}

	file, err := os.Create(savePath + saveName + fileExt)
	if err != nil {
		return err
	}

	writerLen, err := io.Copy(file, bytes.NewReader(body))
	if err != nil {
		return err
	}
	Info("下载总长度：", writerLen)
	return nil
}

/**
获取url返回的内容
*/
func GetByteForUrl(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil || res.StatusCode != 200 {
		return nil, err
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}

//请求url获取文档对象
func Fetch(url string) (*goquery.Document, error) {

	res, err := http.Get(url)
	if err != nil || res.StatusCode != 200 {
		return nil, err
	}

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func FindList(url, find string, callBack func(i int, list *goquery.Selection)) error {
	doc, err := Fetch(url)
	if err != nil {
		return err
	}

	doc.Find(find).Each(func(i int, selection *goquery.Selection) {
		callBack(i, selection)
	})

	return nil
}

func FindOne(url, find string, callBack func(sel *goquery.Selection)) error {
	doc, err := Fetch(url)
	if err != nil {
		return err
	}

	callBack(doc.Find(find))

	return nil
}
