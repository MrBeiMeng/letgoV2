package leetcode_api

import (
	"fmt"
	"letgoV2/system_code/pkg/logging"
	"letgoV2/system_code/service/leetcode_api/leetcode_bodys"
)

type LeetCodeApiImpl struct {
}

func (l *LeetCodeApiImpl) SearchQuestionStatsByTitleSlug(titleSlug string) (error, leetcode_bodys.QuestionStats) {

	postBody := GetSearchQuestionStatsByTitleSlugReqBody(titleSlug)

	err, resp := LeetcodeHttpPost(`https://leetcode.cn/graphql/`, postBody)
	if err != nil {
		return err, leetcode_bodys.QuestionStats{}
	}

	err, questionStats := convQuestionStatsFromResp(resp)
	if err != nil {
		return err, leetcode_bodys.QuestionStats{}
	}

	return nil, questionStats
}

func GetSearchQuestionStatsByTitleSlugReqBody(titleSlug string) string {
	postBody := fmt.Sprintf(`{"query":"\n    query questionStats($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    stats\n  }\n}\n    ","variables":{"titleSlug":"%s"},"operationName":"questionStats"}`, titleSlug)
	return postBody
}

func (l *LeetCodeApiImpl) SearchQuestionTopicTagsByTitleSlug(titleSlug string) (error, []leetcode_bodys.TopicTag) {

	postBody := GetSearchQuestionTopicTagsByTitleSlugReqBody(titleSlug)

	err, resp := LeetcodeHttpPost(`https://leetcode.cn/graphql/`, postBody)
	if err != nil {
		println(err.Error())
	}

	err, topicTags := convQuestionTopicTagsRespFromResp(resp)
	if err != nil {
		return err, nil
	}

	return nil, topicTags
}

func (l *LeetCodeApiImpl) SearchQuestionContentTranslatedByTitleSlug(titleSlug string) (error, leetcode_bodys.TranslatedQuestionContent) {

	postBody := GetSearchQuestionContentTranslatedByTitleSlugReqBody(titleSlug)

	err, resp := LeetcodeHttpPost(`https://leetcode.cn/graphql/`, postBody)
	if err != nil {
		return err, leetcode_bodys.TranslatedQuestionContent{}
	}

	err, questionContent := convTranslatedQuestionContentFromResp(resp)
	if err != nil {
		return err, leetcode_bodys.TranslatedQuestionContent{}
	}

	return nil, questionContent
}

func (l *LeetCodeApiImpl) SearchQuestionContentEnglishByTitleSlug(titleSlug string) (error, string) {

	postBody := GetSearchQuestionContentEnglishByTitleSlugReqBody(titleSlug)

	err, resp := LeetcodeHttpPost(`https://leetcode.cn/graphql/`, postBody)
	if err != nil {
		return err, ""
	}

	err, questionContent := convQuestionContentEnglishFromResp(resp)
	if err != nil {
		return err, ""
	}

	return nil, questionContent
}

func (l *LeetCodeApiImpl) SearchQuestionInfoByTitleSlug(titleSlug string) (error, leetcode_bodys.QuestionInfo) {

	postBody := GetSearchQuestionInfoByTitleSlugReqBody(titleSlug)

	err, resp := LeetcodeHttpPost(`https://leetcode.cn/graphql/`, postBody)
	if err != nil {
		return err, leetcode_bodys.QuestionInfo{}
	}

	err, questionInfo := convQuestionInfoFromResp(resp)
	if err != nil {
		return err, leetcode_bodys.QuestionInfo{}
	}

	return nil, questionInfo
}

func (l *LeetCodeApiImpl) SearchTitleSlugByQuestionId(questionId string) (error, string) {

	// 请求体
	postBody := GetSearchTitleSlugByQuestionIdReqBody(questionId)
	err, resp := LeetcodeHttpPost(`https://leetcode.cn/graphql/`, postBody)
	if err != nil {
		return err, ""
	}

	err, titleSlug := convTitleSlugFromResp(resp)
	if err != nil {
		return err, ""
	}

	logging.Info(fmt.Sprintf("查询到 titleSlug[%s] By questionId[%s]", titleSlug, questionId))
	return nil, titleSlug
}

func (l *LeetCodeApiImpl) SearchGoCodeSnippetByTitleSlug(titleSlug string) (error, string) {

	// 请求体
	postBody := GetSearchGoCodeSnippetByTitleSlugReqBody(titleSlug)

	err, resp := LeetcodeHttpPost(`https://leetcode.cn/graphql/`, postBody)
	if err != nil {
		return err, ""
	}

	err, goCodeSnippet := convGoCodeSnippetFromResp(resp)
	if err != nil {
		return err, ""
	}

	return nil, goCodeSnippet
}
