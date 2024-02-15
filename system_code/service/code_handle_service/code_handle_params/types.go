package code_handle_params

import (
	"fmt"
	"letgoV2/system_code/pkg/util"
	"strings"
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

func (r *RunResult) String() string {
	if r.Pass {
		return r.SuccessSprint()
	}

	return r.FailureSprint()
}

func (r *RunResult) getTestStr() string {

	return strings.ReplaceAll(r.TestStr, "\n", "\\n")
}

func (r *RunResult) SuccessSprint() string {

	return fmt.Sprintf("%-9s | %s | test=%-50s | %s", r.GetDuration(), util.SetColor("PASSED", util.GREEN), util.TruncateString(r.getTestStr(), 50), r.ResultStr)
}

func (r *RunResult) FailureSprint() string {

	return fmt.Sprintf("%s | %s | test=%s | expected=%s | but=%s", r.GetDuration(), util.SetColor("FAIL", util.RED), r.getTestStr(), r.ExpectedResult, r.ResultStr)
}

func (r *RunResult) GetDuration() time.Duration {
	return r.EndTime.Sub(r.StartTime)
}

func NewRunResult(dirId string, resultStr string, success bool, err error, pass bool, testStr string, expectedResult String, startTime, endTime time.Time) RunResult {
	return RunResult{DirId: dirId, ResultStr: resultStr, Success: success, Err: err, Pass: pass, TestStr: testStr, ExpectedResult: expectedResult, StartTime: startTime, EndTime: endTime}
}

type String interface{}

type Test struct {
	TestStr       string `json:"testStr"`
	CorrectResult String `json:"correctResult"`
	ShowWhenErr   string `json:"showWhenErr"`
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
