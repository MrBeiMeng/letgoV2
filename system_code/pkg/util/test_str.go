package util

import (
	"errors"
	"fmt"
	"letgoV2/system_code/pkg/logging"
	"strings"
)

// PeelOffALayer
// 去掉一层外壳 ，例如 func("[[1,2,3],[2,3,4]]") = "[1,2,3]","[2,3,4]"
// 前提是数组或切片类型的字符串参数
func PeelOffALayer(strParams string, trimLeft, trimRight byte) (error, []string) {
	if strParams[0] != trimLeft {
		errStr := fmt.Sprintf("字符串开头非'%c'", trimLeft)
		logging.Info(errStr)
		return errors.New(errStr), nil
	}
	strParams = strParams[1:]

	if strParams[len(strParams)-1] != trimRight {
		errStr := fmt.Sprintf("字符串结尾非'%c'", trimRight)
		logging.Info(errStr)
		return errors.New(errStr), nil
	}
	strParams = strParams[:len(strParams)-1]

	result := make([]string, 0)

	isALayer := !strings.Contains(strParams, "[")
	if isALayer {
		return nil, []string{strParams}
	}

	countDeep := 0
	builder := strings.Builder{}
	for _, subRune := range strParams {
		if countDeep > 0 {
			builder.WriteRune(subRune)
		}
		if subRune == rune(trimLeft) {
			if countDeep > 0 {
				builder.WriteRune(subRune)
			}
			countDeep += 1
			continue
		}
		if subRune == rune(trimRight) {
			countDeep -= 1
			if countDeep == 0 {
				result = append(result, builder.String())
				builder.Reset()
				continue
			}
		}
	}

	if countDeep != 0 {
		errStr := fmt.Sprintf("非法参数:%s", strParams)
		logging.Info(errStr)
		return errors.New(errStr), nil
	}

	return nil, result
}
