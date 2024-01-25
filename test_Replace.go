package main

import (
	"fmt"
	"reflect"
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

func getReplaceMap(value reflect.Value) map[string]string {
	result := make(map[string]string)
	tV := value.Type()

	for i := 0; i < value.NumField(); i++ {
		fieldValue := value.Field(i)
		fieldType := tV.Field(i)
		tag := fieldType.Tag.Get("replace")
		// Check if the field has a "replace" tag
		if tag != "" {
			result[tag] = fieldValue.String()
		}
	}

	return result
}

func main() {

	fileParams := CombinedFileParams{
		GoCodeSnippet:        "sgahsd",
		JsonExampleTestcases: "",
		EnglishTitle:         "",
		Difficulty:           "",
		QuestionId:           "",
		URL:                  "",
		EnglishContent:       "",
		AcRate:               "",
		TranslatedTopicTags:  "",
		TopicTags:            "",
		TranslatedTitle:      "",
		TranslatedContent:    "",
		TitleSlug:            "",
	}
	replaceMap := getReplaceMap(reflect.ValueOf(fileParams))

	fmt.Printf("%+v", replaceMap)
}
