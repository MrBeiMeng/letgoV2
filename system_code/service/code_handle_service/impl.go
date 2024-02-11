package code_handle_service

import (
	"errors"
	"fmt"
	"letgoV2/system_code/pkg/common"
	"letgoV2/system_code/pkg/func_operator"
	"letgoV2/system_code/service/code_handle_service/code_handle_params"
	"reflect"
	"strings"
	"time"
)

type Test struct {
	TestStr       string
	ResultChecker common.CheckDataStruct
	CorrectResult code_handle_params.String
	ShowWhenErr   string
}

type UniqueTests map[string]Test

func (u *UniqueTests) UpsertTests(tests ...code_handle_params.Test) error {
	if (*u) == nil {
		*u = make(UniqueTests)
	}

	for _, test := range tests {
		tempTest := (*u)[test.TestStr]
		tempTest.TestStr = test.TestStr
		tempTest.CorrectResult = test.CorrectResult

		if test.CorrectResult != nil {
			// 对 correctResult 解析
			if str, ok := test.CorrectResult.(string); ok {
				if strings.HasPrefix(str, (&common.MultiSet{}).GetFlag()) {
					err, checker := common.NewMultiSet(str)
					if err != nil {
						return err
					}

					tempTest.ResultChecker = checker
					tempTest.ShowWhenErr = test.ShowWhenErr

					(*u)[test.TestStr] = tempTest
					continue
				}
			}

			checker := common.NewString(test.CorrectResult)

			tempTest.ResultChecker = checker
			tempTest.ShowWhenErr = test.ShowWhenErr
		}

		(*u)[test.TestStr] = tempTest
	}
	return nil
}

type CodeHandleServiceImpl struct {
	FuncMap map[string]interface{}
	TestMap map[string]UniqueTests
}

func (c *CodeHandleServiceImpl) initWhenFirst() {
	if c.FuncMap == nil {
		c.FuncMap = make(map[string]interface{})
	}

	if c.TestMap == nil {
		c.TestMap = make(map[string]UniqueTests)
	}

}

func (c *CodeHandleServiceImpl) updateFunction(dirId string, function interface{}) error {
	c.initWhenFirst()

	if reflect.ValueOf(function).IsNil() {
		return errors.New("nil的函数")
	}
	c.FuncMap[dirId] = function

	return nil
}

func (c *CodeHandleServiceImpl) addTests(dirId string, tests []code_handle_params.Test) error {
	c.initWhenFirst()

	tempTests := c.TestMap[dirId]
	err := tempTests.UpsertTests(tests...)
	if err != nil {
		return err
	}
	c.TestMap[dirId] = tempTests

	return nil
}

func (c *CodeHandleServiceImpl) SignIn(dirId string, function interface{}, testSlice []code_handle_params.Test) error {
	c.initWhenFirst()

	// 进行注册
	if !reflect.ValueOf(function).IsNil() {
		err := c.updateFunction(dirId, function)
		if err != nil {
			return err
		}
	}

	err := c.addTests(dirId, testSlice)
	if err != nil {
		return err
	}

	return nil
}

func (c *CodeHandleServiceImpl) Run(dirId string, test Test, reportChan chan<- code_handle_params.RunResult) time.Time {
	function := c.FuncMap[dirId]
	if reflect.ValueOf(function).IsNil() {
		panic(fmt.Sprintf("未注册方法@%s", dirId))
	}
	startTime := time.Now()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				err := errors.New(fmt.Sprintf("%v", r))
				reportChan <- code_handle_params.NewRunResult(dirId, "", false, err, false, test.TestStr, test.CorrectResult, startTime, time.Now())
			}
		}()

		startTime := time.Now()
		success := true
		err, resultStr := func_operator.RunFunc(function, test.TestStr)
		if err != nil {
			success = false
		}
		endTime := time.Now()

		passed := false
		if test.ResultChecker != nil {
			passed = test.ResultChecker.Check(resultStr)
		}

		reportChan <- code_handle_params.NewRunResult(dirId, resultStr, success, err, passed, test.TestStr, test.CorrectResult, startTime, endTime)
	}()

	return startTime
}

func (c *CodeHandleServiceImpl) AutoRun(dirId string, reportChan chan<- code_handle_params.RunResult) (resultMap map[string]time.Time) {
	resultMap = make(map[string]time.Time)

	tests := c.TestMap[dirId]

	for _, test := range tests {
		resultMap[test.TestStr] = c.Run(dirId, test, reportChan)
	}

	return
}
