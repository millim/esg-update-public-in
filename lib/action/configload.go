package action

import (
	"ali-ecs-public-ip/lib/config"
	"ali-ecs-public-ip/lib/util"
	"errors"
	"io/ioutil"
)

func ConfigLoad() error {
	path := config.YmlPath()
	if path == "" {
		return errors.New("config.yml is null")
	}

	if !util.FileExists(path) {
		return errors.New("no search config file, use -c config.yml")
	}

	ymlFile, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	return config.SetConfig(&ymlFile)
}
