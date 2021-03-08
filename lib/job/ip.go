package job

import (
	"ali-ecs-public-ip/lib/config"
	"io/ioutil"
	"net/http"
)

func getIP() string {
	config := config.GetConfig()
	rep, err := http.Get(config.TestIPURL)
	if err != nil {
		return ""
	}

	defer rep.Body.Close()
	body, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		return ""
	}
	ip := string(body)
	return ip
}

func getOldIP() (string, error) {
	b, err := ioutil.ReadFile(config.PublicIPPath())
	if err != nil {
		println(err.Error())
		return "", err
	}
	return string(b), nil
}
