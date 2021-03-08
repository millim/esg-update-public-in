package job

import (
	"ali-ecs-public-ip/lib/action"
	"ali-ecs-public-ip/lib/config"
	"ali-ecs-public-ip/lib/util"
	"io/ioutil"
	"strings"
	"time"
)

func Run() error {
	for {
		inspect()
		time.Sleep((time.Duration(config.GetConfig().WaitSeconds)) * time.Second)
	}
	return nil
}

func inspect() {
	ip := getIP()
	ip = strings.Replace(ip, " ", "", -1)
	ip = strings.Replace(ip, "\n", "", -1)
	if ip == "" {
		println(time.Now().String(), "获取ip失败")
		return
	}

	if !util.FileExists(config.PublicIPPath()) {
		action.CreateOriginalIP(ip)
		action.UpdateAliIP("", ip)
	} else {
		oldIP, err := getOldIP()
		oldIP = strings.Replace(oldIP, " ", "", -1)
		oldIP = strings.Replace(oldIP, "\n", "", -1)
		if err != nil {
			println(err.Error())
			return
		}
		if oldIP == ip {
			return
		}
		action.UpdateAliIP(oldIP, ip)
		ioutil.WriteFile(config.PublicIPPath(), []byte(ip), 0644)
	}
}
