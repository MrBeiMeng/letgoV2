package generate_service

import (
	"encoding/json"
	"letgoV2/system_code/pkg/logging"
	"letgoV2/system_code/pkg/util"
	"strings"
)

type CombinedFileParams struct {
	GoCodeSnippet        string `replace:"goCodeSnippet"`
	JsonExampleTestcases string
	EnglishTitle         string `replace:"englishTitle"`
	Difficulty           string
	QuestionId           string `replace:"questionId"`
	URL                  string `replace:"url"`
	EnglishContent       string `replace:"englishContent"`
	AcRate               string `replace:"acRate"`
	TranslatedTopicTags  string
	TopicTags            string
	TranslatedTitle      string `replace:"translatedTitle"`
	TranslatedContent    string `replace:"translatedContent"`
	TitleSlug            string
}

type fileCodeParam struct {
	PackageName   string `replace:"packageName"`
	GoCodeSnippet string `replace:"goCodeSnippet"`
}

func newFileCodeParam(packageName string, goCodeSnippet string) *fileCodeParam {
	return &fileCodeParam{PackageName: packageName, GoCodeSnippet: goCodeSnippet}
}

type fileCodeTestParam struct {
	PackageName string `replace:"packageName"`
	FuncName    string `replace:"funcName"`
}

func newFileCodeTestParam(packageName string, goCodeSnippet string) *fileCodeTestParam {

	err, funName := getOneFuncName(goCodeSnippet)
	if err != nil {
		logging.Error(err)
	}
	return &fileCodeTestParam{PackageName: packageName, FuncName: funName}
}

type FileMetaData struct {
	PackageName string `replace:"packageName"`
	SampleTests string `replace:"sampleTests"`
}

func NewFileMetaData(packageName string, jsonExampleTestcases string) *FileMetaData {
	args := make([]string, 0)

	err := json.Unmarshal([]byte(jsonExampleTestcases), &args)
	if err != nil {
		logging.Error(err)
	}
	//for i := range args {
	//	args[i] = fmt.Sprintf("\"%s\"", args)
	//}

	return &FileMetaData{PackageName: packageName, SampleTests: strings.Join(args, ",")}
}

type fileReadMeEnParam struct {
	EnglishTitle   string `replace:"englishTitle"`
	DifficultySpan string `replace:"difficultySpan"`
	QuestionId     string `replace:"questionId"`
	URL            string `replace:"url"`
	EnglishContent string `replace:"englishContent"`
	AcRate         string `replace:"acRate"`
	TagList        string `replace:"tagList"`
}

func newFileReadMeEnParam(param CombinedFileParams) *fileReadMeEnParam {
	f := fileReadMeEnParam{}
	util.CopyStructFields(param, &f)
	f.DifficultySpan = getEnDifficultySpan(getDifficulty(param.Difficulty))
	f.TagList = param.TopicTags

	return &f
}

type fileReadMeZhParam struct {
	TranslatedTitle   string `replace:"translatedTitle"`
	DifficultySpan    string `replace:"difficultySpan"`
	QuestionId        string `replace:"questionId"`
	URL               string `replace:"url"`
	TranslatedContent string `replace:"translatedContent"`
	AcRate            string `replace:"acRate"`
	TagList           string `replace:"tagList"`
}

func newFileReadMeZhParam(param CombinedFileParams) *fileReadMeZhParam {
	f := fileReadMeZhParam{}
	util.CopyStructFields(param, &f)
	f.DifficultySpan = getZhDifficultySpan(getDifficulty(param.Difficulty))
	f.TagList = param.TranslatedTopicTags

	return &f
}
