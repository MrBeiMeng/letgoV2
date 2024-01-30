package code_handle_service

import (
	"errors"
	"fmt"
	"letgoV2/system_code/pkg/func_operator"
	"letgoV2/system_code/service/code_handle_service/code_handle_params"
	"time"
)

type CodeHandleServiceImpl struct {
	FuncTests map[string]struct {
		Function interface{}
		TS       code_handle_params.TestSlice
	}
}

func (c *CodeHandleServiceImpl) SignIn(dirId string, function interface{}, testSlice []code_handle_params.Test) error {
	if c.FuncTests == nil {
		c.FuncTests = make(map[string]struct {
			Function interface{}
			TS       code_handle_params.TestSlice
		})
	}

	if _, ok := c.FuncTests[dirId]; ok {
		fT := c.FuncTests[dirId]
		fT.TS.Merge(testSlice)

		if function != nil {
			fT.Function = function
		}
		c.FuncTests[dirId] = fT
	} else {
		slice := code_handle_params.NewTestSlice()
		slice.Merge(testSlice)

		c.FuncTests[dirId] = struct {
			Function interface{}
			TS       code_handle_params.TestSlice
		}{Function: function, TS: *slice}
	}

	return nil
}

func (c *CodeHandleServiceImpl) GetTests(dirId string) []code_handle_params.Test {
	ts := c.FuncTests[dirId].TS
	return ts.GetTests()
}

func (c *CodeHandleServiceImpl) Run(dirId string, test code_handle_params.Test, reportChan chan<- code_handle_params.RunResult) time.Time {
	function := c.FuncTests[dirId].Function
	startTime := time.Now()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				err := errors.New(fmt.Sprintf("%v", r))
				reportChan <- code_handle_params.NewRunResult(dirId, "", false, err, test)
			}
		}()

		success := true
		err, resultStr := func_operator.RunFunc(function, test.TestStr)
		if err != nil {
			success = false
		}

		reportChan <- code_handle_params.NewRunResult(dirId, resultStr, success, err, test)
	}()

	return startTime
}

func (c *CodeHandleServiceImpl) AutoRun(dirId string, reportChan chan<- code_handle_params.RunResult) (resultMap map[string]time.Time) {
	resultMap = make(map[string]time.Time)

	s := c.FuncTests[dirId]
	tests := s.TS.GetTests()

	for _, test := range tests {
		resultMap[test.TestStr] = c.Run(dirId, test, reportChan)
	}

	return
}
