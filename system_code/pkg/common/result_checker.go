package common

import (
	"errors"
	"strings"
)

type AlwaysPass struct {
}

func NewAlwaysPass() *AlwaysPass {
	return &AlwaysPass{}
}

func (a *AlwaysPass) Check(checkStr string) bool {
	return true
}

type String struct {
	val   string
	isNil bool
}

func NewString(val interface{}) *String {
	obj := &String{}
	if val != nil {
		obj.val = val.(string)
		obj.isNil = false
	} else {
		obj.isNil = true
	}

	return obj
}

func (s *String) Check(checkStr string) bool {
	return s.isNil || strings.EqualFold(s.val, checkStr)
}

type MultiSet struct {
	dataSet map[string]int
}

func (m *MultiSet) GetFlag() string {
	return "@multiset"
}

func NewMultiSet(expectedResult string) (error, *MultiSet) {
	// expectedResult like "@multiset[1,2]"
	flag := (&MultiSet{}).GetFlag()
	if !strings.HasPrefix(expectedResult, flag) {
		return errors.New("don't have prefix '@multiset'"), nil
	}

	dataStr := strings.ReplaceAll(expectedResult, flag, "")
	splitter := NewSplitUtil([]byte(dataStr))
	elements := splitter.GetLevelElements(0, true)

	dataSet := make(map[string]int)
	for _, element := range elements {
		num := 1
		if count, ok := dataSet[element]; ok {
			num = count + 1
		}

		dataSet[element] = num
	}

	result := &MultiSet{}
	result.dataSet = dataSet
	return nil, result
}

func (m *MultiSet) Check(result string) bool {
	splitter := NewSplitUtil([]byte(result))
	elements := splitter.GetLevelElements(0, true)

	tempMap := make(map[string]int)

	for key, value := range m.dataSet {
		tempMap[key] = value
	}

	for _, element := range elements {
		if count, ok := tempMap[element]; ok && count > 0 {
			tempMap[element] = count - 1
		} else {

			return false
		}

	}

	for _, value := range tempMap {
		if value > 0 {
			return false
		}
	}

	return true
}
