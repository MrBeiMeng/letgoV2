package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"letgoV2/system_code/commands/cmd_params"
	"letgoV2/system_code/pkg/logging"
	"letgoV2/system_code/service/code_handle_service"
	"letgoV2/system_code/service/code_handle_service/code_handle_params"
	"strings"

	_ "letgoV2/your_code"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "letgo run -d zzzz ",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if runParam.DirId != "" {
			dirId := strings.TrimSpace(runParam.DirId)
			reportChan := make(chan code_handle_params.RunResult, 100)

			//tests := code_handle_service.CodeHandleService.GetTests(strings.TrimSpace(runParam.DirId))
			startMap := code_handle_service.CodeHandleService.AutoRun(dirId, reportChan)
			logging.Info(fmt.Sprintf("共有%d个test等待测试", len(startMap)))
			countDone := 0

			for {
				select {
				case r := <-reportChan:

					fmt.Printf("%+v\n", strings.ReplaceAll(fmt.Sprintf("%+v", r), "\n", "\\n"))
					countDone += 1
				default:
					fmt.Printf("loading \n")
				}
				if len(startMap) == countDone {
					logging.Info(fmt.Sprintf("程序结束 len = %d countDone = %d", len(startMap), countDone))
					break
				}
			}
		}

	},
}

var runParam cmd_params.RunParam

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&runParam.DirId, "dir", "d", "", "letgo run -d zzzz")
}
