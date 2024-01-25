package main

import (
	"encoding/json"
	"fmt"
	. "letgoV2/system_code/service/leetcode_api"
	"letgoV2/system_code/service/leetcode_api/leetcode_bodys"
)

func main() {

	postBody := fmt.Sprintf(`{"query":"\n    query questionTranslations($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    translatedTitle\n    translatedContent\n  }\n}\n    ","variables":{"titleSlug":"%s"},"operationName":"questionTranslations"}`, "median-of-two-sorted-arrays")

	err, resp := LeetcodeHttpPost(`https://leetcode.cn/graphql/`, postBody)
	if err != nil {
		println(err.Error())
	}

	var translatedQuestionContentResp leetcode_bodys.TranslatedQuestionContentResp
	var translatedQuestionContent leetcode_bodys.TranslatedQuestionContent
	err = json.Unmarshal(resp, &translatedQuestionContentResp)
	if err != nil {
		println(err.Error())
	}

	translatedQuestionContent = translatedQuestionContentResp.Data.QuestionContent

	fmt.Printf("%#v\n", translatedQuestionContent)

	fmt.Println(translatedQuestionContent.Content)

}
