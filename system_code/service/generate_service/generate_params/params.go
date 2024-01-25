package generate_params

import (
	"encoding/json"
	"fmt"
)

type GenerateParams struct {
	QuestionInfo QuestionInfo
	MetaData     MetaData
	Tests        []string
	Code         string
}

func GetGenerateParams(questionId string, titleSlug string, metaData string, tests string, code string) (GenerateParams, error) {

	data, err := GetMetaData(metaData)
	if err != nil {
		return GenerateParams{}, err
	}

	testArr, err := getTests(tests)
	if err != nil {
		return GenerateParams{}, err
	}
	return GenerateParams{
		QuestionInfo: GetQuestionInfo(questionId, titleSlug),
		MetaData:     data,
		Tests:        testArr,
		Code:         code,
	}, nil
}

func getTests(tests string) ([]string, error) {

	testArr := make([]string, 0)

	err := json.Unmarshal([]byte(tests), &testArr)
	if err != nil {

		return nil, err
	}

	return testArr, nil
}

type QuestionInfo struct {
	Id        string
	TitleSlug string
	Url       string
}

func GetQuestionInfo(questionId string, titleSlug string) QuestionInfo {

	url := fmt.Sprintf("https://leetcode.cn/problems/%s", titleSlug)

	return QuestionInfo{
		Id:        questionId,
		TitleSlug: titleSlug,
		Url:       url,
	}
}

type MetaData struct {
	Name   string
	Params []struct {
		Name string
		Type string
	}
	Return struct {
		Type string
	}
}

func GetMetaData(metaData string) (MetaData, error) {
	resultMetaData := MetaData{}

	err := json.Unmarshal([]byte(metaData), &resultMetaData)
	if err != nil {
		return resultMetaData, err
	}

	return resultMetaData, nil
}
