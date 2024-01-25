package main

import (
	"encoding/json"
	"fmt"
	. "letgoV2/system_code/service/leetcode_api"
	"letgoV2/system_code/service/leetcode_api/leetcode_bodys"
	"strings"
)

func main() {

	postBody := fmt.Sprintf(`{"query":"\n    query questionStats($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    stats\n  }\n}\n    ","variables":{"titleSlug":"%s"},"operationName":"questionStats"}`, "median-of-two-sorted-arrays")

	err, resp := LeetcodeHttpPost(`https://leetcode.cn/graphql/`, postBody)
	if err != nil {
		println(err.Error())
	}

	var questionStatsResp leetcode_bodys.QuestionStatsResp
	var questionStats leetcode_bodys.QuestionStats
	err = json.Unmarshal(resp, &questionStatsResp)
	if err != nil {
		println(err.Error())
	}

	err = json.Unmarshal([]byte(strings.ReplaceAll(questionStatsResp.Data.Question.Stats, "\\", "")), &questionStats)
	if err != nil {
		println(err.Error())
	}

	fmt.Printf("%+v", questionStats)

}
