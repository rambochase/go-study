package error

import (
	"strings"
	"bytes"
	"fmt"
)

type ErrorType string

//和对应模块的错误常量
const (
	//下载器错误
	ERROR_TYPE_DOWNLOADER ErrorType = "download error"
	//分析器错误
	ERROR_TYPE_ANALYZER ErrorType = "analyzer error"
	//条目处理管道错误
	ERROR_TYPE_PIPELINE ErrorType = "pipeline error"
	//调度器错误
	ERROR_TYPE_SCHEDULER ErrorType = "scheduler error"
)
//错误类型
type CrawlerError interface {
	Type() ErrorType//获取错误的类型
	//获取错误提示信息
	Error() string
}

type myCrawlerError struct {
	errType ErrorType
	errMsg string
	fullErrMsg string
}

func CreateCrawlerError(errType ErrorType,errMsg string) CrawlerError {
	return &myCrawlerError{errType:errType,errMsg:strings.TrimSpace(errMsg)}
}

func (e *myCrawlerError) Type() ErrorType {
	return e.errType
}

func (e *myCrawlerError) Error() string {
	if e.fullErrMsg == "" {
		e.genFullErrMsg()
	}
	return e.fullErrMsg
}

func (e *myCrawlerError) genFullErrMsg() {
	var buffer bytes.Buffer
	buffer.WriteString("crawler error: ")
	if e.errType != "" {
		buffer.WriteString(string(e.errType))
		buffer.WriteString(": ")
	}
	buffer.WriteString(e.errMsg)
	e.fullErrMsg = fmt.Sprintf("%s",buffer.String())
	return 
}