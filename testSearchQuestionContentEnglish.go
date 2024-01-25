package main

import (
	"encoding/json"
	"fmt"
	. "letgoV2/system_code/service/leetcode_api"
	"letgoV2/system_code/service/leetcode_api/leetcode_bodys"
)

func main() {
	postBody := fmt.Sprintf(`{"query":"\n    query questionContent($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    content\n    editorType\n    mysqlSchemas\n    dataSchemas\n  }\n}\n    ","variables":{"titleSlug":"%s"},"operationName":"questionContent"}`, "median-of-two-sorted-arrays")

	err, resp := LeetcodeHttpPost(`https://leetcode.cn/graphql/`, postBody)
	if err != nil {
		println(err.Error())
	}

	var dataContentResp leetcode_bodys.QuestionContentResp
	var questionContent string
	err = json.Unmarshal(resp, &dataContentResp)
	if err != nil {
		println(err.Error())
	}

	questionContent = dataContentResp.Data.Question.Content

	println(questionContent)
}
