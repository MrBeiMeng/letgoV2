package down_service

import (
	"fmt"
	"letgoV2/system_code/service/generate_service"
)

// 整合API得到的信息到预生成files中
func convQuestionInfoToGParam(info QuestionInfoMerged) (result generate_service.CombinedFileParams) {
	result.GoCodeSnippet = info.GoCodeSnippet
	result.JsonExampleTestcases = info.JsonExampleTestcases
	result.EnglishTitle = info.Title
	result.Difficulty = info.Difficulty
	result.QuestionId = info.QuestionId
	result.EnglishContent = info.EnglishContent
	result.AcRate = info.AcRate
	result.TopicTags = info.TopicTags
	result.TranslatedTopicTags = info.TranslatedTopicTags
	result.TranslatedTitle = info.TranslatedTitle
	result.TranslatedContent = info.TranslatedContent
	result.TitleSlug = info.TitleSlug
	result.URL = fmt.Sprintf("https://leetcode.cn/problems/%s", result.TitleSlug)

	return
}
