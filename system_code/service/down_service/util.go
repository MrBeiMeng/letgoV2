package down_service

import (
	"errors"
	"fmt"
	"letgoV2/system_code/pkg/logging"
	"letgoV2/system_code/pkg/util"
	"letgoV2/system_code/service/leetcode_api"
	"letgoV2/system_code/service/leetcode_api/leetcode_bodys"
	"strings"
)

func getMergedQuestionInfo(titleSlug string) (questionInfo QuestionInfoMerged, err error) {
	apiService := leetcode_api.LeetCodeApi

	err, snippet := apiService.SearchGoCodeSnippetByTitleSlug(titleSlug)
	if err != nil {
		return
	}
	questionInfo.GoCodeSnippet = snippet
	err, info := apiService.SearchQuestionInfoByTitleSlug(titleSlug)
	if err != nil {
		return
	}
	util.CopyStructFields(info, &questionInfo)
	err, enContent := apiService.SearchQuestionContentEnglishByTitleSlug(titleSlug)
	if err != nil {
		return
	}
	questionInfo.EnglishContent = enContent
	err, contentStruct := apiService.SearchQuestionContentTranslatedByTitleSlug(titleSlug)
	if err != nil {
		return
	}
	questionInfo.TranslatedTitle = contentStruct.Title
	questionInfo.TranslatedContent = contentStruct.Content
	err, tags := apiService.SearchQuestionTopicTagsByTitleSlug(titleSlug)
	if err != nil {
		return
	}
	trTags := make([]string, 0)
	enTags := make([]string, 0)
	for _, tagStruct := range tags {
		trTags = append(trTags, tagStruct.TranslatedName)
		enTags = append(enTags, tagStruct.Name)
	}
	questionInfo.TranslatedTopicTags = strings.Join(trTags, "|")
	questionInfo.TopicTags = strings.Join(enTags, "|")
	err, stats := apiService.SearchQuestionStatsByTitleSlug(titleSlug)
	if err != nil {
		return
	}
	util.CopyStructFields(stats, &questionInfo)
	err, testInfo := apiService.SearchQuestionSampleTestInfoByTitleSlug(titleSlug)
	if err != nil {
		return
	}
	util.CopyStructFields(testInfo, &questionInfo)

	return
}
func getGoCodeTemplate(graphqlResp leetcode_bodys.GraphqlResp) (leetcode_bodys.CodeSnippet, error) {
	// --- 获取golang代码模板
	var golangCodeSnippet leetcode_bodys.CodeSnippet

	for _, snippet := range graphqlResp.Data.Question.CodeSnippets {
		if snippet.LangSlug == "golang" {
			golangCodeSnippet = snippet
		}
	}

	if golangCodeSnippet.LangSlug == "" {
		errStr := fmt.Sprintf("not have LangSlug = golang")
		logging.Info(errStr)
		return leetcode_bodys.CodeSnippet{}, errors.New(errStr)
	}
	// --- 获取golang代码模板
	return golangCodeSnippet, nil
}
