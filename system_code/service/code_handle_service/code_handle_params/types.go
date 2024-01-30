package code_handle_params

type RunResult struct {
	DirId     string
	ResultStr string
	Success   bool // 是否运行成功
	Err       error
	Test
}

func NewRunResult(dirId string, resultStr string, success bool, err error, test Test) RunResult {
	return RunResult{DirId: dirId, ResultStr: resultStr, Success: success, Err: err, Test: test}
}

type String interface{}

type Test struct {
	TestStr       string
	CorrectResult String
	ShowWhenErr   string
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
