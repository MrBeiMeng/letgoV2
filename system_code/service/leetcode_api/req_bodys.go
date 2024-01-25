package leetcode_api

import "fmt"

func GetSearchQuestionTopicTagsByTitleSlugReqBody(titleSlug string) string {
	postBody := fmt.Sprintf(`{"query":"\n    query singleQuestionTopicTags($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    topicTags {\n      name\n      slug\n      translatedName\n    }\n  }\n}\n    ","variables":{"titleSlug":"%s"},"operationName":"singleQuestionTopicTags"}`, titleSlug)
	return postBody
}

func GetSearchQuestionContentTranslatedByTitleSlugReqBody(titleSlug string) string {
	postBody := fmt.Sprintf(`{"query":"\n    query questionTranslations($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    translatedTitle\n    translatedContent\n  }\n}\n    ","variables":{"titleSlug":"%s"},"operationName":"questionTranslations"}`, titleSlug)
	return postBody
}

func GetSearchQuestionContentEnglishByTitleSlugReqBody(titleSlug string) string {
	postBody := fmt.Sprintf(`{"query":"\n    query questionContent($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    content\n    editorType\n    mysqlSchemas\n    dataSchemas\n  }\n}\n    ","variables":{"titleSlug":"%s"},"operationName":"questionContent"}`, titleSlug)
	return postBody
}

func GetSearchQuestionInfoByTitleSlugReqBody(titleSlug string) string {
	postBody := fmt.Sprintf(`{"query":"\n    query questionTitle($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    questionId\n    questionFrontendId\n    title\n    titleSlug\n    isPaidOnly\n    difficulty\n    likes\n    dislikes\n    categoryTitle\n  }\n}\n    ","variables":{"titleSlug":"%s"},"operationName":"questionTitle"}`, titleSlug)
	return postBody
}

func GetSearchTitleSlugByQuestionIdReqBody(questionId string) string {
	postBody := fmt.Sprintf(`{"query":"\n    query latestUserQuestionRecordsAndRemarks($questionId: ID!) {\n  latestUserRecords: latestUserRecordByQuestionId(questionId: $questionId) {\n    dailyQuestion {\n      date\n      id\n      question {\n        titleSlug\n        title\n      }\n      userStatus\n    }\n    isDailyQuestion\n    weeklyQuestion {\n      date\n      id\n      weekNum\n      question {\n        titleSlug\n        title\n      }\n      userStatus\n    }\n    isWeeklyQuestion\n  }\n  remarks: userDailyQuestionRemarkCards {\n    cardCount\n    usedCount\n  }\n}\n    ","variables":{"questionId":"%s"},"operationName":"latestUserQuestionRecordsAndRemarks"}`, questionId)
	return postBody
}

func GetSearchGoCodeSnippetByTitleSlugReqBody(titleSlug string) string {
	postBody := fmt.Sprintf(`{"query":"\n    query questionEditorData($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    questionId\n    questionFrontendId\n    codeSnippets {\n      lang\n      langSlug\n      code\n    }\n    envInfo\n    enableRunCode\n    hasFrontendPreview\n    frontendPreviews\n  }\n}\n    ","variables":{"titleSlug":"%s"},"operationName":"questionEditorData"}`, titleSlug)
	return postBody
}

func GetSearchQuestionSampleTestInfoByTitleSlugReqBody(titleSlug string) string {
	postBody := fmt.Sprintf(`{"query":"\n    query consolePanelConfig($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    questionId\n    questionFrontendId\n    questionTitle\n    enableRunCode\n    enableSubmit\n    enableTestMode\n    jsonExampleTestcases\n    exampleTestcases\n    metaData\n    sampleTestCase\n  }\n}\n    ","variables":{"titleSlug":"%s"},"operationName":"consolePanelConfig"}`, titleSlug)
	return postBody
}
