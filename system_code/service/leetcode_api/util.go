package leetcode_api

import "letgoV2/system_code/pkg/util/config_util"

func getHeaderMap() map[string]string {

	fieldLeetcode := config_util.Fields("Leetcode")
	headerMap := make(map[string]string)
	headerMap["content-type"] = fieldLeetcode.Get("ContentType")
	headerMap["origin"] = fieldLeetcode.Get("Origin")
	headerMap["user-agent"] = fieldLeetcode.Get("UserAgent")

	return headerMap
}
