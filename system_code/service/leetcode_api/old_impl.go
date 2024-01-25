package leetcode_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"letgoV2/system_code/pkg/logging"
	"letgoV2/system_code/pkg/setting"
	"letgoV2/system_code/pkg/util"
	"letgoV2/system_code/service/leetcode_api/leetcode_bodys"
)

func (l *LeetCodeApiImpl) SearchQuestionByQuestionId(questionId string) (error, leetcode_bodys.GraphqlResp) {
	err, titleSlug := l.SearchTitleSlugByQuestionId(questionId)
	var graphqlResp leetcode_bodys.GraphqlResp
	if err != nil {
		return err, graphqlResp
	}

	if titleSlug == "" {
		errStr := fmt.Sprintf("did not found titleSlug")
		logging.Info(errStr)
		return errors.New(errStr), graphqlResp
	}
	logging.Info(fmt.Sprintf("got titleSlug [%s] by questionID [%s]", titleSlug, questionId))

	err, resp := l.SearchQuestionByTitleSlug(titleSlug)
	if err != nil {
		return err, resp
	}

	return nil, resp
}

func (l *LeetCodeApiImpl) SearchQuestionSampleTestInfoByTitleSlug(titleSlug string) (error, leetcode_bodys.QuestionSampleTestInfo) {
	// 获取列表
	postBody := GetSearchQuestionSampleTestInfoByTitleSlugReqBody(titleSlug)
	err, resp := LeetcodeHttpPost(`https://leetcode.cn/graphql/`, postBody)
	if err != nil {
		return nil, leetcode_bodys.QuestionSampleTestInfo{}
	}

	var questionInfoResp leetcode_bodys.QuestionSampleTestInfoResp

	err = json.Unmarshal(resp, &questionInfoResp)
	if err != nil {
		errStr := fmt.Sprintf("JSON 解析错误:%s", err)
		logging.Info(errStr)
		return errors.New(errStr), questionInfoResp.Data.Question
	}

	return nil, questionInfoResp.Data.Question
}

func (l *LeetCodeApiImpl) SearchQuestionByTitleSlug(titleSlug string) (error, leetcode_bodys.GraphqlResp) {
	// 模拟请求
	headerMap := setting.LeetCodeConf.HeaderMap
	cookies := setting.LeetCodeConf.Cookies

	// 请求体
	postBody := fmt.Sprintf(`{"query":"\n    query questionEditorData($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    questionId\n    questionFrontendId\n    codeSnippets {\n      lang\n      langSlug\n      code\n    }\n    envInfo\n    enableRunCode\n    hasFrontendPreview\n    frontendPreviews\n  }\n}\n    ","variables":{"titleSlug":"%s"},"operationName":"questionEditorData"}`, titleSlug)

	err, resp := util.HttpPost(`https://leetcode.cn/graphql/`, cookies, headerMap, postBody)
	if err != nil {
		return nil, leetcode_bodys.GraphqlResp{}
	}

	var graphqlResp leetcode_bodys.GraphqlResp

	err = json.Unmarshal(resp, &graphqlResp)
	if err != nil {
		errStr := fmt.Sprintf("JSON 解析错误:%s", err)
		logging.Info(errStr)

		return errors.New(errStr), graphqlResp
	}

	return nil, graphqlResp
}
