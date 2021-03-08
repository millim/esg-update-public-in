package main

import (
	"ali-ecs-public-ip/lib/action"
	"ali-ecs-public-ip/lib/config"
	"ali-ecs-public-ip/lib/job"
)

func main() {
	config.FlagInit()

	action.PidOut()

	err := action.ConfigLoad()
	if err != nil {
		panic(err)
		return
	}

	err = job.Run()
	if err != nil {
		panic(err)
	}
}
