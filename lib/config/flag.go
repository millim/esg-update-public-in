package config

import "flag"

var publicIPPath string
var pidfile string
var configYmlPath string

func FlagInit() {
	flag.StringVar(&configYmlPath, "c", "", "-c config.yml")
	flag.StringVar(&publicIPPath, "f", "", "-f public_ip")
	flag.StringVar(&pidfile, "p", "", "-p pid")
	flag.Parse()
}

func PidFilePath() string {
	return pidfile
}

func PublicIPPath() string {
	return publicIPPath
}

func YmlPath() string {
	return configYmlPath
}
