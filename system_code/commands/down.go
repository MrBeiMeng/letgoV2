package commands

import (
	"fmt"
	"letgoV2/system_code/commands/cmd_params"
	"letgoV2/system_code/pkg/logging"
	"letgoV2/system_code/service/down_service"

	"github.com/spf13/cobra"
)

// downCmd represents the down command
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "将目标题目提取到本地 /your_code/下 例如: letgo down -i 88",
	Long:  `将目标题目提取到本地 /your_code/下 例如: letgo down -i 88`,
	Run: func(cmd *cobra.Command, args []string) {
		callingMethod(cmd.Use)

		downService := down_service.DownService

		// 条件判断执行哪个指令
		if downParam.ID != "" {
			logging.Info(fmt.Sprintf("letgo down func called --id [%s]", downParam.ID))

			err := downService.DownById(downParam.ID)
			if err != nil {
				logging.Info(err.Error())
			}

			return
		}

		if downParam.TitleSlug != "" {
			logging.Info(fmt.Sprintf("letgo down func called --title [%s]", downParam.TitleSlug))

			err := downService.DownByTitleSlug(downParam.TitleSlug)
			if err != nil {
				logging.Info(err.Error())
			}

			return
		}

		logging.Info("no args error")
	},
}

var downParam cmd_params.DownParam

func init() {
	rootCmd.AddCommand(downCmd)
	downCmd.Flags().StringVarP(&downParam.ID, "id", "i", "", "letgo down -i 88")
	downCmd.Flags().StringVarP(&downParam.TitleSlug, "title", "t", "", "letgo down -t longest-substring-without-repeating-characters")
}
