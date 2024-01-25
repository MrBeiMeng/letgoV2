package main

import (
	"letgoV2/system_code/commands"
	"letgoV2/system_code/pkg/logging"
)

func main() {

	logging.Info("-------- start --------")
	commands.Execute()
	logging.Info("-------- end --------")
}
