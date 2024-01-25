package down_service

import (
	"errors"
	"fmt"
	"letgoV2/system_code/pkg/logging"
	"letgoV2/system_code/service/generate_service"
	"letgoV2/system_code/service/leetcode_api"
)

type DownServiceImpl struct {
}

func (d *DownServiceImpl) DownByTitleSlug(titleSlug string) (err error) {
	if titleSlug == "" {
		errStr := fmt.Sprintf("title slug is empty")
		logging.Info(errStr)
		return errors.New(errStr)
	}

	questionInfo, err := getMergedQuestionInfo(titleSlug) // 合并所有api的信息

	combinedFileParams := convQuestionInfoToGParam(questionInfo)

	err = generate_service.GenerateService.GenerateFiles(combinedFileParams)
	if err != nil {
		return
	}

	return nil
}

func (d *DownServiceImpl) DownById(questionId string) error {
	if questionId == "" {
		errStr := fmt.Sprintf("questionId is empty")
		logging.Info(errStr)
		return errors.New(errStr)
	}

	err, titleSlug := leetcode_api.LeetCodeApi.SearchTitleSlugByQuestionId(questionId)
	if err != nil {
		return err
	}

	// 调用方法
	return d.DownByTitleSlug(titleSlug)
}
