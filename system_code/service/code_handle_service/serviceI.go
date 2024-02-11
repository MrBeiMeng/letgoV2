package code_handle_service

import (
	"letgoV2/system_code/service/code_handle_service/code_handle_params"
	"time"
)

type CodeHandleServiceI interface {
	// SignIn 如果已注册，则第二次注册时function参数是可选的。否则会返回异常
	SignIn(dirId string, function interface{}, testStrSlice []code_handle_params.Test) error
	Run(dirId string, test Test, reportChan chan<- code_handle_params.RunResult) time.Time
	// AutoRun map[string]time.Time 保存了key(dirId)对应的(value)开始运行时间
	AutoRun(dirId string, reportChan chan<- code_handle_params.RunResult) map[string]time.Time
}

var CodeHandleService CodeHandleServiceI = &CodeHandleServiceImpl{}
