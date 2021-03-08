package action

import (
	"ali-ecs-public-ip/lib/config"
	"fmt"
	"io/ioutil"
	"os"
)

func PidOut() {
	if config.PidFilePath() != "" {
		ioutil.WriteFile(config.PidFilePath(), []byte(fmt.Sprintf("%d", os.Getpid())), 0644)
	}
}
