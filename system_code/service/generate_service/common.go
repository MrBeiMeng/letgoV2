package generate_service

import (
	"fmt"
	"letgoV2/system_code/pkg/logging"
	"strings"
)

type difficulty int

const (
	EASY difficulty = iota
	MEDIUM
	HARD
)

var (
	zhDifficultyFlags = []string{
		`<span style="background:#3de1ad;border-radius:5px;padding:1px 5px;font-weight:bold;color:#ffffff">简单</span> `,
		`<span style="background:#ffa400;border-radius:5px;padding:1px 5px;font-weight:bold;color:#ffffff">中等</span> `,
		`<span style="background:#f35336;border-radius:5px;padding:1px 5px;font-weight:bold;color:#ffffff">困难</span>`,
	}

	enDifficultyFlags = []string{
		`<span style="background:#3de1ad;border-radius:5px;padding:1px 5px;font-weight:bold;color:#ffffff">EASY</span> `,
		`<span style="background:#ffa400;border-radius:5px;padding:1px 5px;font-weight:bold;color:#ffffff">MEDIUM</span> `,
		`<span style="background:#f35336;border-radius:5px;padding:1px 5px;font-weight:bold;color:#ffffff">HARD</span>`,
	}
)

func getDifficulty(difficultyStr string) difficulty {
	switch strings.ToLower(difficultyStr) {
	case "easy":
		return EASY
	case "medium":
		return MEDIUM
	case "hard":
		return HARD
	default:
		logging.Error(fmt.Sprintf("未捕获到的难度级别:[%s]", difficultyStr))
	}

	return -1
}

func getZhDifficultySpan(difficulty difficulty) string {
	return zhDifficultyFlags[difficulty]
}

func getEnDifficultySpan(difficulty difficulty) string {
	return enDifficultyFlags[difficulty]
}
