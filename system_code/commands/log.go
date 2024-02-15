package commands

import (
	"fmt"
	"letgoV2/system_code/pkg/logging"
)

func callingMethod(name string) {
	logging.Info(fmt.Sprintf("letgo %s calling", name))
}
