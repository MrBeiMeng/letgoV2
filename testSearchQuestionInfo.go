package main

import (
	"encoding/json"
	"fmt"
	. "letgoV2/system_code/service/leetcode_api"
	"letgoV2/system_code/service/leetcode_api/leetcode_bodys"
)

func main() {

	postBody := fmt.Sprintf(`{"query":"\n    query questionTitle($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    questionId\n    questionFrontendId\n    title\n    titleSlug\n    isPaidOnly\n    difficulty\n    likes\n    dislikes\n    categoryTitle\n  }\n}\n    ","variables":{"titleSlug":"%s"},"operationName":"questionTitle"}`, "median-of-two-sorted-arrays")

	err, resp := LeetcodeHttpPost(`https://leetcode.cn/graphql/`, postBody)
	if err != nil {

		println(err.Error())
	}

	var questionInfoResp leetcode_bodys.QuestionInfoResp

	err = json.Unmarshal(resp, &questionInfoResp)
	if err != nil {
		println(err.Error())
	}

	fmt.Printf("%v", questionInfoResp)

}
