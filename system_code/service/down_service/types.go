package down_service

type QuestionInfoMerged struct {
	GoCodeSnippet        string
	QuestionId           string `json:"questionId"`
	QuestionFrontendId   string `json:"questionFrontendId"`
	Title                string `json:"title"`
	TitleSlug            string
	IsPaidOnly           bool   `json:"isPaidOnly"`
	Difficulty           string `json:"difficulty"`
	Likes                int    `json:"likes"`
	Dislikes             int    `json:"dislikes"`
	CategoryTitle        string `json:"categoryTitle"`
	EnglishContent       string
	TranslatedTitle      string `json:"translatedTitle"`
	TranslatedContent    string `json:"translatedContent"`
	TopicTags            string // 通过 | 分隔
	TranslatedTopicTags  string
	TotalAccepted        string `json:"totalAccepted"`
	TotalSubmission      string `json:"totalSubmission"`
	TotalAcceptedRaw     int    `json:"totalAcceptedRaw"`
	TotalSubmissionRaw   int    `json:"totalSubmissionRaw"`
	AcRate               string `json:"acRate"`
	QuestionTitle        string `json:"questionTitle"`
	EnableRunCode        bool   `json:"enableRunCode"`
	EnableSubmit         bool   `json:"enableSubmit"`
	EnableTestMode       bool   `json:"enableTestMode"`
	JsonExampleTestcases string `json:"jsonExampleTestcases"`
	ExampleTestcases     string `json:"exampleTestcases"`
	MetaData             string `json:"metaData"`
	SampleTestCase       string `json:"sampleTestCase"`
}
