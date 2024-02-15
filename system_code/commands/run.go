package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"letgoV2/system_code/commands/cmd_params"
	"letgoV2/system_code/pkg/logging"
	"letgoV2/system_code/pkg/util"
	"letgoV2/system_code/service/code_handle_service"
	"letgoV2/system_code/service/code_handle_service/code_handle_params"
	_ "letgoV2/your_code"
	"strings"
	"time"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "letgo run -d zzzz ",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		callingMethod(cmd.Use)
		defer func() {
			if r := recover(); r != nil {
				logging.Error(r)
			}
		}()

		if runParam.DirId != "" {
			dirId := strings.TrimSpace(runParam.DirId)
			reportChan := make(chan code_handle_params.RunResult, 100)

			//tests := code_handle_service.CodeHandleService.GetTests(strings.TrimSpace(runParam.DirId))
			startMap := code_handle_service.CodeHandleService.AutoRun(dirId, reportChan)
			logging.Info(fmt.Sprintf("共有%d个test等待测试", len(startMap)))
			countDone := 0

			running := true

			for running {
				select {
				case r := <-reportChan:
					fmt.Printf("\r")

					if r.Err != nil {
						logging.Error(r.Err)
						running = false
						break
					}

					if !r.Pass {
						running = false
					}

					logging.Info(r.String())

					countDone += 1
				default:
					waiting(&count)
					time.Sleep(100 * time.Millisecond)
				}
				if len(startMap) == countDone {
					logging.Info(fmt.Sprintf("程序结束 共 %d 完成 %d %s", len(startMap), countDone, util.SetColor("🎉 ALL PASSED!!", util.GREEN)))
					break
				}
			}
		}

	},
}

var (
	flags = []string{"—", "\\", "|", "/"}
	count = 1
)

func waiting(count *int) {
	fmt.Printf("\r%s", flags[(*count)%len(flags)])
	*count += 1
}

var runParam cmd_params.RunParam

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&runParam.DirId, "dir", "d", "", "letgo run -d zzzz")
}
