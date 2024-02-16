package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"letgoV2/system_code/commands/cmd_params"
	"letgoV2/system_code/pkg/logging"
	"letgoV2/system_code/pkg/util/config_util"
	"strings"
)

// downCmd represents the down command
var setCmd = &cobra.Command{
	Use:   "config",
	Short: "设置参数到配置文件，请不要上传你的配置文件 system_code/conf",
	Run: func(cmd *cobra.Command, args []string) {
		if setParam.Set != "" {
			setValue := setParam.Set

			if strings.Contains(setValue, "=") {
				dataList := strings.Split(setValue, "=")
				key, value := dataList[0], dataList[1]
				key = strings.TrimSpace(key)

				setUtil := config_util.Config
				subSetUtil := setUtil.GetData()
				datas := strings.Split(key, ".")
				for i, subFix := range datas {

					if i == len(datas)-1 {
						subSetUtil.Set(subFix, value)
						break
					}

					subSetUtil = config_util.Fields(subFix)
				}

				value = strings.TrimSpace(value)

				if !removeParam.Confirm {
					logging.Warn(fmt.Sprintf("将更新配置文件 set key[%s]=value[%s] [Y/n]", key, value))
					input := ""

					fmt.Scanln(&input)
					if !strings.Contains(strings.ToLower(strings.TrimSpace(input)), "y") {

						logging.Info("用户取消")

						return
					}

				}

				setUtil.Save()
			}
		} else if setParam.Show != "" {
			fmt.Printf("\"%s\":\"%s\"\n", setParam.Show, config_util.Get(setParam.Show))
		} else if setParam.Cookies != "" {
			config_util.Fields("Leetcode").Set("Cookies", setParam.Cookies)
			config_util.Save()
		}
	},
}

var setParam cmd_params.SetParam

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.Flags().StringVar(&setParam.Set, "set", "", "letgo config set a=\"b\"")
	setCmd.Flags().StringVar(&setParam.Show, "show", "", "letgo config show")
	setCmd.Flags().StringVar(&setParam.Cookies, "cookies", "", "letgo config show")
	setCmd.PersistentFlags().BoolVarP(&setParam.Confirm, "yes", "y", false, "-y 表示确认")
}
