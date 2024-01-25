package main

import (
	"encoding/json"
	"fmt"
	"letgoV2/system_code/pkg/util"
	"letgoV2/system_code/service/leetcode_api/leetcode_bodys"
)

func main() {

	// 模拟请求

	headerMap := make(map[string]string)

	cookies := "gr_user_id=388ae01d-3b2e-4de4-a070-8e152d41d98f; _bl_uid=wvlgzrmIhg5n4nfw8cb6lpy003e8; csrftoken=OPjaqhfiy22TmuqK7nz9AoXt4gAKjrKR1WcM8wFvCpFnZAauffNStbptLkAKdzy0; LEETCODE_SESSION=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJfYXV0aF91c2VyX2lkIjoiNDA3MDI5OCIsIl9hdXRoX3VzZXJfYmFja2VuZCI6ImRqYW5nby5jb250cmliLmF1dGguYmFja2VuZHMuTW9kZWxCYWNrZW5kIiwiX2F1dGhfdXNlcl9oYXNoIjoiZDVjNWJhYmUzYjZmODQxYzYwN2ViNGIxMGZhNjY0N2Y1MmEwMGZmZmRhNDhlZTlmNDdhZDhkOGMwNGJmNDY3NiIsImlkIjo0MDcwMjk4LCJlbWFpbCI6IjExOTIzODQ3MjJAcXEuY29tIiwidXNlcm5hbWUiOiJiZWltZW5nY2x1YiIsInVzZXJfc2x1ZyI6ImJlaW1lbmdjbHViIiwiYXZhdGFyIjoiaHR0cHM6Ly9hc3NldHMubGVldGNvZGUuY24vYWxpeXVuLWxjLXVwbG9hZC91c2Vycy9iZWltZW5nY2x1Yi9hdmF0YXJfMTY2NzUzNTc0OC5wbmciLCJwaG9uZV92ZXJpZmllZCI6dHJ1ZSwiX3RpbWVzdGFtcCI6MTcwNTQ4NzkzMi44Nzk2MTI0LCJleHBpcmVkX3RpbWVfIjoxNzA4MDIzNjAwLCJ2ZXJzaW9uX2tleV8iOjB9.Nc4sHmn7JNWgNtoMSD90WD-_NSp1iPx5-BJPQhlSmu4; a2873925c34ecbd2_gr_last_sent_cs1=beimengclub; Hm_lvt_f0faad39bcf8471e3ab3ef70125152c3=1704936546,1705487245,1705720037; a2873925c34ecbd2_gr_session_id=dad5eb27-ec82-4bf5-b4eb-af6b4eb357ce; a2873925c34ecbd2_gr_last_sent_sid_with_cs1=dad5eb27-ec82-4bf5-b4eb-af6b4eb357ce; a2873925c34ecbd2_gr_session_id_sent_vst=dad5eb27-ec82-4bf5-b4eb-af6b4eb357ce; _gid=GA1.2.1486669979.1705720037; _gat=1; _ga=GA1.1.1165414047.1704936546; Hm_lpvt_f0faad39bcf8471e3ab3ef70125152c3=1705720813; a2873925c34ecbd2_gr_cs1=beimengclub; _ga_PDVPZYN3CW=GS1.1.1705720036.5.1.1705720820.45.0.0"
	headerMap["content-type"] = "application/json"
	headerMap["origin"] = "https://leetcode.cn"
	headerMap["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36"

	// 获取列表

	postBody := fmt.Sprintf(`{"query":"\n    query questionEditorData($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    questionId\n    questionFrontendId\n    codeSnippets {\n      lang\n      langSlug\n      code\n    }\n    envInfo\n    enableRunCode\n    hasFrontendPreview\n    frontendPreviews\n  }\n}\n    ","variables":{"titleSlug":"%s"},"operationName":"questionEditorData"}`, "longest-substring-without-repeating-characters")
	//postBody := "{\"query\":\"\\n    query problemsetQuestionList($categorySlug: String, $limit: Int, $skip: Int, $filters: QuestionListFilterInput) {\\n  problemsetQuestionList(\\n    categorySlug: $categorySlug\\n    limit: $limit\\n    skip: $skip\\n    filters: $filters\\n  ) {\\n    hasMore\\n    total\\n    questions {\\n      acRate\\n      difficulty\\n      freqBar\\n      frontendQuestionId\\n      isFavor\\n      paidOnly\\n      solutionNum\\n      status\\n      title\\n      titleCn\\n      titleSlug\\n      topicTags {\\n        name\\n        nameTranslated\\n        id\\n        slug\\n      }\\n      extra {\\n        hasVideoSolution\\n        topCompanyTags {\\n          imgUrl\\n          slug\\n          numSubscribed\\n        }\\n      }\\n    }\\n  }\\n}\\n    \",\"variables\":{\"categorySlug\":\"\",\"skip\":%d,\"limit\":100,\"filters\":{}}}"
	_, resp := util.HttpPost(`https://leetcode.cn/graphql/`, cookies, headerMap, postBody)

	var graphqlResp leetcode_bodys.GraphqlResp

	err := json.Unmarshal(resp, &graphqlResp)
	if err != nil {
		fmt.Println("JSON 解析错误:", err)
		return
	}

	fmt.Printf("%v", graphqlResp)

	// 现在，你可以访问解析后的数据，例如：
	fmt.Println("QuestionContent ID:", graphqlResp.Data.Question.QuestionId)
	fmt.Println("QuestionContent Frontend ID:", graphqlResp.Data.Question.QuestionFrontendId)

	// 遍历 CodeSnippets
	for _, snippet := range graphqlResp.Data.Question.CodeSnippets {
		fmt.Println("Lang:", snippet.Lang)
		fmt.Println("LangSlug:", snippet.LangSlug)
		fmt.Println("Code:", snippet.Code)
		fmt.Println("--------")
	}

}
