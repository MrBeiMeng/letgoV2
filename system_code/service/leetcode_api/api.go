package leetcode_api

import (
	"letgoV2/system_code/service/leetcode_api/leetcode_bodys"
)

type LeetCodeApiI interface {
	SearchTitleSlugByQuestionId(questionId string) (error, string)

	// 新方法
	SearchGoCodeSnippetByTitleSlug(titleSlug string) (error, string)                                               // 代码模板
	SearchQuestionInfoByTitleSlug(titleSlug string) (error, leetcode_bodys.QuestionInfo)                           // ！难度/题目类别
	SearchQuestionContentEnglishByTitleSlug(titleSlug string) (error, string)                                      // 题目描述
	SearchQuestionContentTranslatedByTitleSlug(titleSlug string) (error, leetcode_bodys.TranslatedQuestionContent) // 中文题目/中文描述
	SearchQuestionTopicTagsByTitleSlug(titleSlug string) (error, []leetcode_bodys.TopicTag)                        // 题目标签
	SearchQuestionStatsByTitleSlug(titleSlug string) (error, leetcode_bodys.QuestionStats)                         // 题目信息统计 "{\"totalAccepted\": \"1.1M\", \"totalSubmission\": \"2.5M\", \"totalAcceptedRaw\": 1059264, \"totalSubmissionRaw\": 2527818, \"acRate\": \"41.9%\"}"

	// 弃用方法
	//SearchQuestionByTitleSlug(titleSlug string) (error, leetcode_bodys.GraphqlResp)   // 获取代码模板 √已重写
	//SearchQuestionByQuestionId(questionId string) (error, leetcode_bodys.GraphqlResp) // 避免使用，即将弃用

	// 旧方法，但是还在用
	SearchQuestionSampleTestInfoByTitleSlug(titleSlug string) (error, leetcode_bodys.QuestionSampleTestInfo) // 简单测试
}

var LeetCodeApi LeetCodeApiI

func init() {
	LeetCodeApi = &LeetCodeApiImpl{}
}
