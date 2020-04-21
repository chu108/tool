package tool

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

const (
	POST = "POST"
	GET  = "GET"
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

	filePath := savePath + saveName + fileExt
	if IsExist(filePath) {
		Err(saveName, "文件已存在")
		return nil
	}

	body, err := GetByteForUrl(url)
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
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

//http get请求
func Get(url string, data map[string]interface{}, header map[string]string) ([]byte, error) {
	return Request(GET, url, data, header)
}

//http post请求
func Post(url string, data map[string]interface{}, header map[string]string) ([]byte, error) {
	return Request(POST, url, data, header)
}

func Request(method, requestUrl string, data map[string]interface{}, header map[string]string) ([]byte, error) {
	method = strings.ToUpper(method)
	client := &http.Client{
		Timeout:   3 * time.Second, //超时为3秒
		Transport: http.DefaultTransport,
	}
	var (
		req  *http.Request
		err  error
		body io.Reader = nil
	)

	dataLen := len(data)
	switch method {
	case POST:
		if dataLen > 0 {
			bytesData, err := json.Marshal(data)
			if err != nil {
				return nil, err
			}
			body = bytes.NewReader(bytesData)
		}
		req, err = http.NewRequest(POST, requestUrl, body)
	case GET:
		if dataLen > 0 {
			params := url.Values{}
			for key, val := range data {
				if value, ok := val.(string); ok {
					params.Add(key, value)
				}
				if value, ok := val.(int); ok {
					params.Add(key, strconv.Itoa(value))
				}
			}
			URL, err := url.Parse(requestUrl)
			if err != nil {
				return nil, err
			}
			URL.RawQuery = params.Encode()
			requestUrl = URL.String()
		}
		req, err = http.NewRequest(GET, requestUrl, nil)
	}
	if err != nil {
		return nil, err
	}

	headerLen := len(header)
	if headerLen > 0 {
		for key, val := range header {
			if key == "cookie" {
				SetCookie(req, val)
			}
			req.Header.Add(key, val)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

//设置cookie
func SetCookie(req *http.Request, cookie string) {
	cookie = strings.TrimSpace(cookie)
	cks := strings.Split(cookie, ";")
	for _, v := range cks {
		item := strings.Split(v, "=")
		cookieItem := &http.Cookie{Name: item[0], Value: item[1]}
		req.AddCookie(cookieItem)
	}
}

//写入文件
func WriteFile(filePath string, body []byte) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	writerLen, err := io.Copy(file, bytes.NewReader(body))
	if err != nil {
		return err
	}
	Info("写入长度：", writerLen)
	return nil
}

//tcp端口检测
func TcpGather(ip, port string) bool {
	if ip == "" || port == "" {
		return false
	}
	addr := net.JoinHostPort(ip, port)
	timeout := 3 * time.Second
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return false
	}
	defer conn.Close()
	if conn != nil {
		return true
	}
	return false
}
