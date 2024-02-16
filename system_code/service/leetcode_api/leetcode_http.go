package leetcode_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"letgoV2/system_code/pkg/setting"
	"letgoV2/system_code/pkg/util"
	"strings"
)

type ErrResp struct {
	Errors []ErrorBody `json:"errors"`
}

type ErrorBody struct {
	Message   interface{} `json:"message"`
	Locations interface{} `json:"locations"`
	Path      interface{} `json:"path"`
}

func LeetcodeHttpPost(url string, reqBody string) (error, []byte) {

	// 模拟请求
	headerMap := setting.LeetCodeConf.HeaderMap
	cookies := setting.LeetCodeConf.Cookies

	err, bytes := util.HttpPost(url, cookies, headerMap, reqBody)
	if err != nil {
		return err, nil
	}

	errResp := ErrResp{}

	err = json.Unmarshal(bytes, &errResp)
	if err != nil {
		return err, nil
	}

	if errResp.Errors != nil || len(errResp.Errors) != 0 {
		errStr := strings.Builder{}

		for _, errBody := range errResp.Errors {
			errStr.WriteString(fmt.Sprintf("msg:%v|Locations:%v|Path:%v", errBody.Message, errBody.Locations, errBody.Path))
		}

		return errors.New(errStr.String()), nil
	}

	return err, bytes
}
