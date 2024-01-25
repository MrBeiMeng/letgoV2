package leetcode_api

import (
	"letgoV2/system_code/pkg/setting"
	"letgoV2/system_code/pkg/util"
)

func LeetcodeHttpPost(url string, reqBody string) (error, []byte) {

	// 模拟请求
	headerMap := setting.LeetCodeConf.HeaderMap
	cookies := setting.LeetCodeConf.Cookies

	return util.HttpPost(url, cookies, headerMap, reqBody)
}
