package main

import (
	"encoding/json"
	"fmt"
	. "letgoV2/system_code/service/leetcode_api"
	"letgoV2/system_code/service/leetcode_api/leetcode_bodys"
)

func main() {
	postBody := fmt.Sprintf(`{"query":"\n    query singleQuestionTopicTags($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    topicTags {\n      name\n      slug\n      translatedName\n    }\n  }\n}\n    ","variables":{"titleSlug":"%s"},"operationName":"singleQuestionTopicTags"}`, "median-of-two-sorted-arrays")

	err, resp := LeetcodeHttpPost(`https://leetcode.cn/graphql/`, postBody)
	if err != nil {
		println(err.Error())
	}

	var topicTags []leetcode_bodys.TopicTag
	var questionTopicTagsResp leetcode_bodys.QuestionTopicTagsResp
	err = json.Unmarshal(resp, &questionTopicTagsResp)
	if err != nil {
		println(err.Error())
	}

	topicTags = questionTopicTagsResp.Data.Question.TopicTags

	fmt.Printf("%+v", topicTags)

}
