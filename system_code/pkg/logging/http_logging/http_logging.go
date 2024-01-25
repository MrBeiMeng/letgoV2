package http_logging

import (
	"errors"
	"fmt"
	"io/ioutil"
	"letgoV2/system_code/pkg/setting"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var (
	F2                 *os.File
	DefaultPrefix      = ""
	DefaultCallerDepth = 2
	LogSaveName        = "log"
	LogFileExt         = "log"
	TimeFormat         = "20060102"

	httpLogger      *log.Logger
	httpLogPrefix   = ""
	HttpLogSavePath = "http_logs/"
)

func getHttpLogFilePath() string {

	logDir := setting.LogDir

	fullLogSavePath := path.Join(logDir, HttpLogSavePath)

	return fmt.Sprintf("%s/", fullLogSavePath)
}

func GetLogFileFullPath(prefixPath string) string {
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)

	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func mkDir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+getHttpLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func OpenLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("Permission :%v", err)
	}

	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}

	return handle
}

func init() {
	filePath2 := GetLogFileFullPath(getHttpLogFilePath())
	F2 = OpenLogFile(filePath2)

	httpLogger = log.New(F2, DefaultPrefix, log.LstdFlags)
}

func HttpPostLog(client *http.Client, req *http.Request) ([]byte, error) {
	startTime := time.Now()
	// 执行代码段
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	setHttpPrefix(req.Method, time.Now().Sub(startTime).String())

	// ---- 日志打印
	err, respBodyStr := printHttpLog(req, resp)
	if err != nil {
		return nil, err
	}
	// ---- 日志打印

	return respBodyStr, err
}

func printHttpLog(req *http.Request, resp *http.Response) (error, []byte) {
	// 请求url
	targetUrl := req.URL
	cookies := make([]string, 0)
	for _, cookie := range req.Cookies() {
		cookies = append(cookies, cookie.String())
	}
	//cookiesStr := strings.Join(cookies, "; ")
	headerMap := (map[string][]string)(req.Header)
	headerStr := strings.Builder{}
	for key, value := range headerMap {
		headerStr.WriteString(key + ":" + strings.Join(value, ","))
	}

	// 结果记录
	//reqestBodyStr, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	panic(err)
	//}

	// 结果记录
	respBodyStr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	//logStr := fmt.Sprintf("url: %s | resp: %s | header: %s | cookies : %s | req: %s", targetUrl, string(respBodyStr), headerStr, cookiesStr, string(reqestBodyStr))
	logStr := fmt.Sprintf("| url: %s | status: %s | resp: %s ", targetUrl, resp.Status, string(respBodyStr))

	httpLogger.Println(logStr)

	// 返回状态检测
	if resp.StatusCode != 200 {
		errStr := fmt.Sprintf("访问异常: [%s]", resp.Status)
		return errors.New(errStr), nil
	}

	return nil, respBodyStr
}

func setHttpPrefix(method string, duration string) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		httpLogPrefix = fmt.Sprintf("  %s  |  %s:%d  |  %s  |", strings.ToUpper(method), filepath.Base(file), line, duration)
	} else {
		httpLogPrefix = fmt.Sprintf("  %s  |  %s  |", strings.ToUpper(method), duration)
	}

	httpLogger.SetPrefix(httpLogPrefix)
}
