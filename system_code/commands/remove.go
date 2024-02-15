package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"letgoV2/system_code/commands/cmd_params"
	"letgoV2/system_code/pkg/logging"
	"letgoV2/system_code/pkg/util"
	"os"
	"path/filepath"
	"strings"
)

// downCmd represents the down command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "删掉不想要的文件夹 例如: letgo remove -d zzzy",
	Run: func(cmd *cobra.Command, args []string) {
		callingMethod(cmd.Use)

		if removeParam.DirId != "" {
			if strings.Contains(removeParam.DirId, "zzzz") {
				logging.Warn("零号文件夹仅用来示例，请不要删除。可能导致错误")
				return
			}

			dirPath := util.SearchDir(removeParam.DirId, "your_code")

			if !removeParam.Confirm {
				logging.Warn(fmt.Sprintf("将删除%s文件夹以及其所有文件 [Y/n]", dirPath))
				input := ""

				fmt.Scanln(&input)
				if !strings.Contains(strings.ToLower(strings.TrimSpace(input)), "y") {

					logging.Info("用户取消")

					return
				}

			}

			// 寻找包含dirId 的文件夹
			// 删除所有文件夹
			// 删除enter.go 中的行

			err := os.RemoveAll(dirPath)
			if err != nil {
				logging.Error(err)
				return
			}

			logging.Warn(fmt.Sprintf("已删除%s", dirPath))

			dir, _ := os.Getwd()
			enterGoPath := filepath.Join(dir, "your_code/enter.go")
			bytes, err := os.ReadFile(enterGoPath)
			if err != nil {
				logging.Error(err)
			}

			str := fmt.Sprintf("%s", bytes)

			lines := strings.Split(str, "\n")

			file, err := os.OpenFile(enterGoPath, os.O_RDWR, os.ModePerm)
			if err != nil {
				panic(err)
			}

			file.Truncate(0)
			for _, line := range lines {
				if !strings.Contains(line, removeParam.DirId) && line != "\n" {
					file.Write([]byte(fmt.Sprintf("%s\n", line)))
				}
			}

			logging.Warn(fmt.Sprintf("已重写%s", enterGoPath))

		}

	},
}

var removeParam cmd_params.RemoveParam

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.Flags().StringVarP(&removeParam.DirId, "dir", "d", "", "letgo remove -d zzzy")
	removeCmd.PersistentFlags().BoolVarP(&removeParam.Confirm, "yes", "y", false, "letgo remove -d zzzy -y")
}
