package leetcode_bodys

type TopicTag struct {
	Name           string `json:"name"`
	Slug           string `json:"slug"`
	TranslatedName string `json:"translatedName"`
}

type CodeSnippet struct {
	Lang     string `json:"lang"`
	LangSlug string `json:"langSlug"`
	Code     string `json:"code"`
}

type Question struct {
	QuestionId         string        `json:"questionId"`
	QuestionFrontendId string        `json:"questionFrontendId"`
	CodeSnippets       []CodeSnippet `json:"codeSnippets"`
}

// UserRecords 定义结构体以映射JSON数据
type UserRecords struct {
	Data struct {
		LatestUserRecords struct {
			DailyQuestion []struct {
				Date     string `json:"date"`
				ID       string `json:"id"`
				Question struct {
					TitleSlug string `json:"titleSlug"`
					Title     string `json:"title"`
				} `json:"question"`
				UserStatus string `json:"userStatus"`
			} `json:"dailyQuestion"`
			IsDailyQuestion  bool          `json:"isDailyQuestion"`
			WeeklyQuestion   []interface{} `json:"weeklyQuestion"`
			IsWeeklyQuestion bool          `json:"isWeeklyQuestion"`
		} `json:"latestUserRecords"`
		Remarks struct {
			CardCount int `json:"cardCount"`
			UsedCount int `json:"usedCount"`
		} `json:"remarks"`
	} `json:"data"`
}

// QuestionSampleTestInfo 定义结构体以匹配JSON结构
type QuestionSampleTestInfo struct {
	QuestionID           string `json:"questionId"`
	QuestionFrontendID   string `json:"questionFrontendId"`
	QuestionTitle        string `json:"questionTitle"`
	EnableRunCode        bool   `json:"enableRunCode"`
	EnableSubmit         bool   `json:"enableSubmit"`
	EnableTestMode       bool   `json:"enableTestMode"`
	JsonExampleTestcases string `json:"jsonExampleTestcases"`
	ExampleTestcases     string `json:"exampleTestcases"`
	MetaData             string `json:"metaData"`
	SampleTestCase       string `json:"sampleTestCase"`
}

type TranslatedQuestionContent struct {
	Title   string `json:"translatedTitle"`
	Content string `json:"translatedContent"`
}

type QuestionStats struct {
	TotalAccepted      string `json:"totalAccepted"`
	TotalSubmission    string `json:"totalSubmission"`
	TotalAcceptedRaw   int    `json:"totalAcceptedRaw"`
	TotalSubmissionRaw int    `json:"totalSubmissionRaw"`
	AcRate             string `json:"acRate"`
}

type QuestionInfo struct {
	QuestionId         string `json:"questionId"`
	QuestionFrontendId string `json:"questionFrontendId"`
	Title              string `json:"title"`
	TitleSlug          string `json:"titleSlug"`
	IsPaidOnly         bool   `json:"isPaidOnly"`
	Difficulty         string `json:"difficulty"`
	Likes              int    `json:"likes"`
	Dislikes           int    `json:"dislikes"`
	CategoryTitle      string `json:"categoryTitle"`
}
