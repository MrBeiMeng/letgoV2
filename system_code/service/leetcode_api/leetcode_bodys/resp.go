package leetcode_bodys

type QuestionTopicTagsResp struct {
	Data struct {
		Question struct {
			TopicTags []TopicTag `json:"topicTags"`
		} `json:"question"`
	} `json:"data"`
}

type GraphqlResp struct {
	Data struct {
		Question Question `json:"question"`
	} `json:"data"`
}

type TranslatedQuestionContentResp struct {
	Data struct {
		QuestionContent TranslatedQuestionContent `json:"question"`
	} `json:"data"`
}

type QuestionStatsResp struct {
	Data struct {
		Question struct {
			Stats string `json:"stats"`
		} `json:"question"`
	} `json:"data"`
}

type QuestionInfoResp struct {
	Data struct {
		QuestionInfo QuestionInfo `json:"question"`
	} `json:"data"`
}

type QuestionSampleTestInfoResp struct {
	Data struct {
		Question QuestionSampleTestInfo `json:"question"`
	} `json:"data"`
}

type QuestionContentResp struct {
	Data struct {
		Question struct {
			Content      string        `json:"content"`
			EditorType   string        `json:"editorType"`
			MysqlSchemas []interface{} `json:"mysqlSchemas"`
			DataSchemas  []interface{} `json:"dataSchemas"`
		} `json:"question"`
	} `json:"data"`
}
