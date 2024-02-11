package code_handle_params

import (
	"time"
)

type RunResult struct {
	DirId          string
	ResultStr      string
	Success        bool // 是否运行成功
	Err            error
	Pass           bool
	TestStr        string
	ExpectedResult String
	StartTime      time.Time
	EndTime        time.Time
}

func (r *RunResult) GetDuration() time.Duration {
	return r.EndTime.Sub(r.StartTime)
}

func NewRunResult(dirId string, resultStr string, success bool, err error, pass bool, testStr string, expectedResult String, startTime, endTime time.Time) RunResult {
	return RunResult{DirId: dirId, ResultStr: resultStr, Success: success, Err: err, Pass: pass, TestStr: testStr, ExpectedResult: expectedResult, StartTime: startTime, EndTime: endTime}
}

type String interface{}

type Test struct {
	TestStr       string
	CorrectResult String
	ShowWhenErr   string
}

type UniqueTests struct {
	NameMap map[string]Test
}

type TestSlice struct {
	NameMap map[string]Test
}

func NewTestSlice() *TestSlice {
	return &TestSlice{NameMap: make(map[string]Test)}
}

func (s *TestSlice) Merge(add []Test) {
	for _, test := range add {
		if _, ok := s.NameMap[test.TestStr]; !ok {
			s.NameMap[test.TestStr] = test
			continue
		}

		if test.CorrectResult != nil {
			s.NameMap[test.TestStr] = test
		}

	}
}

func (s *TestSlice) GetTests() (result []Test) {
	for _, value := range s.NameMap {
		result = append(result, value)
	}

	return
}
