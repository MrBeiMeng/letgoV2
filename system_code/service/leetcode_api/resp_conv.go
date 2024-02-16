package leetcode_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"letgoV2/system_code/pkg/logging"
	"letgoV2/system_code/service/leetcode_api/leetcode_bodys"
	"strings"
)

func convQuestionStatsFromResp(resp []byte) (error, leetcode_bodys.QuestionStats) {
	var questionStatsResp leetcode_bodys.QuestionStatsResp
	var questionStats leetcode_bodys.QuestionStats
	err := json.Unmarshal(resp, &questionStatsResp)
	if err != nil {
		return err, leetcode_bodys.QuestionStats{}
	}

	err = json.Unmarshal([]byte(strings.ReplaceAll(questionStatsResp.Data.Question.Stats, "\\", "")), &questionStats)
	if err != nil {
		return err, leetcode_bodys.QuestionStats{}
	}
	return nil, questionStats
}

// convQuestionTopicTagsRespFromResp
func convQuestionTopicTagsRespFromResp(resp []byte) (error, []leetcode_bodys.TopicTag) {

	var questionTopicTagsResp leetcode_bodys.QuestionTopicTagsResp
	err := json.Unmarshal(resp, &questionTopicTagsResp)
	if err != nil {
		return err, nil
	}

	return nil, questionTopicTagsResp.Data.Question.TopicTags
}

func convTranslatedQuestionContentFromResp(resp []byte) (error, leetcode_bodys.TranslatedQuestionContent) {
	var translatedQuestionContentResp leetcode_bodys.TranslatedQuestionContentResp
	err := json.Unmarshal(resp, &translatedQuestionContentResp)
	if err != nil {
		return err, leetcode_bodys.TranslatedQuestionContent{}
	}

	return nil, translatedQuestionContentResp.Data.QuestionContent
}

func convQuestionContentEnglishFromResp(resp []byte) (error, string) {
	var dataContentResp leetcode_bodys.QuestionContentResp
	err := json.Unmarshal(resp, &dataContentResp)
	if err != nil {
		return err, ""
	}

	return nil, dataContentResp.Data.Question.Content
}

func convQuestionInfoFromResp(resp []byte) (error, leetcode_bodys.QuestionInfo) {
	var questionInfoResp leetcode_bodys.QuestionInfoResp
	var questionInfo leetcode_bodys.QuestionInfo

	err := json.Unmarshal(resp, &questionInfoResp)
	if err != nil {
		return err, questionInfo
	}
	questionInfo = questionInfoResp.Data.QuestionInfo
	return nil, questionInfo
}

func convGoCodeSnippetFromResp(resp []byte) (error, string) {
	var graphqlResp leetcode_bodys.GraphqlResp

	err := json.Unmarshal(resp, &graphqlResp)
	if err != nil {
		errStr := fmt.Sprintf("JSON 解析错误:%s", err)
		logging.Info(errStr)

		return errors.New(errStr), ""
	}

	// --- 获取golang代码模板
	var golangCodeSnippet leetcode_bodys.CodeSnippet

	for _, snippet := range graphqlResp.Data.Question.CodeSnippets {
		if snippet.LangSlug == "golang" {
			golangCodeSnippet = snippet
		}
	}

	if golangCodeSnippet.LangSlug == "" {
		errStr := fmt.Sprintf("未在查询结果中找到go代码模板")
		logging.Error(errStr)
		return errors.New(errStr), ""
	}
	// --- 获取golang代码模板
	return nil, golangCodeSnippet.Code
}

func convTitleSlugFromResp(resp []byte) (error, string) {
	var userRecordsResp leetcode_bodys.UserRecords

	err := json.Unmarshal(resp, &userRecordsResp)
	if err != nil {
		errStr := fmt.Sprintf("JSON 解析错误:[%s] | json: [%s]", err, string(resp))
		logging.Info(errStr)

		return errors.New(errStr), ""
	}

	if len(userRecordsResp.Data.LatestUserRecords.DailyQuestion) <= 0 {
		errStr := fmt.Sprintf("userRecordsResp.Data.LatestUserRecords.DailyQuestion.len = 0 这个接口影响的代码非常多，意味着你要改很多代码了。[:)]")

		logging.Error(errStr)
		return errors.New(errStr), ""

	}

	return nil, userRecordsResp.Data.LatestUserRecords.DailyQuestion[0].Question.TitleSlug
}
