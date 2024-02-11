package util

import (
	"fmt"
	"reflect"
	"strings"
)

type SplitUtilI interface {
	// GetLevelElements level from 0
	GetLevelElements(level int, withChild bool) []string
	LevelNum() (min, max int)
}

type SplitUtil struct {
	data []interface{}
}

func iterateSlice(minLevel, maxLevel *int, slice []interface{}) {

	// 如果所有元素都是数组类型，则minLevel加一
	// 有一个是数组类型，maxLevel加一
	hasSubArr, allSubArr := false, true

	for i := 0; i < len(slice); i++ {
		if value, ok := slice[i].([]interface{}); ok {
			iterateSlice(minLevel, maxLevel, value)
			hasSubArr = true
		} else {
			allSubArr = false
		}
	}

	if hasSubArr {
		*maxLevel += 1
	}

	if allSubArr {
		*minLevel += 1
	}

}

func (s *SplitUtil) LevelNum() (min, max int) {

	iterateSlice(&min, &max, s.data)

	return min, max
}

func sliceLevelElements(targetLevel, currentLevel int, slice []interface{}, withChild bool, result *[]string) {

	for i := 0; i < len(slice); i++ {
		tempValue := slice[i]
		if currentLevel < targetLevel {
			if value, ok := tempValue.([]interface{}); ok {
				sliceLevelElements(targetLevel, currentLevel+1, value, withChild, result)
			}

			continue
		}

		if _, ok := tempValue.([]interface{}); !ok || withChild {
			if reflect.TypeOf(tempValue).Kind() == reflect.Uint8 {
				*result = append(*result, fmt.Sprintf("%c", tempValue))
			} else {
				*result = append(*result, fmt.Sprintf("%s", tempValue))
			}
		}

	}

}

func (s *SplitUtil) GetLevelElements(level int, withChild bool) []string {
	tempLevelElements := make([]string, 0)

	sliceLevelElements(level, 0, s.data, withChild, &tempLevelElements)

	return tempLevelElements
}

func isSwapBySquareBrackets(str string) bool {
	if len(str) < 2 {
		return false
	}

	return str[0] == '[' && str[len(str)-1] == ']'
}

func test1(strData *[]byte) (result []interface{}) {

	// 先检测是否有特殊标识
	if (*strData)[0] == '[' {
		*strData = (*strData)[1:] // 去除外面括号
	}

	for len(*strData) > 0 {
		temChar := (*strData)[0]
		if temChar == '@' {
			tempElement := strings.Builder{}
			for len(*strData) > 0 {
				subTempChar := (*strData)[0]

				tempElement.WriteByte(subTempChar)

				*strData = (*strData)[1:]
				if subTempChar == ']' {
					result = append(result, tempElement.String())
					break
				}

			}

			continue
		}

		if temChar == '[' {

			result = append(result, test1(strData))
			continue
		}

		if !strings.Contains("[],", string(temChar)) {
			result = append(result, temChar)
		}

		*strData = (*strData)[1:]
		if temChar == ']' {
			break
		}
	}

	return result
}

// NewSplitUtil 这个类负责拆分字符串，像是[[1,2,3],2,3]
// 将他们拆分成层级
// 如果任意层级包含@则不进行拆分
// 前提：[]是分组标识，, 是间隔字符
func NewSplitUtil(strData []byte) *SplitUtil {
	obj := &SplitUtil{}
	obj.data = make([]interface{}, 0)

	obj.data = test1(&strData)

	//// 先检测是否有特殊标识
	//if isSwapBySquareBrackets(string(strData)) {
	//	strData = strData[1 : len(strData)-1] // 去除外面括号
	//}
	//
	//for len(strData) > 0 {
	//	tempChar := strData[0]
	//	if tempChar == '[' {
	//		obj.data = append(obj.data, test1(&strData))
	//		continue
	//	}
	//	if !strings.Contains("[],", string(tempChar)) {
	//		obj.data = append(obj.data, tempChar)
	//	}
	//
	//	strData = strData[1:]
	//
	//	if tempChar == ']' {
	//		break
	//	}
	//}

	return obj
}
