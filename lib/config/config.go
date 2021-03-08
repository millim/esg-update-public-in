package config

import "gopkg.in/yaml.v2"

type Group struct {
	GroupId string `yaml:"groupId"`
	Port    string `yaml:"port"`
	Info    string `yaml:"info"`
}

type Config struct {
	TestIPURL       string  `yaml:"testUrl"`
	GetUpdateURL    string  `yaml:"updateUrl"`
	WaitSeconds     int64   `yaml:"waitSeconds"`
	RegionID        string  `yaml:"regionId"`
	AccessKeyID     string  `yaml:"accessKeyId"`
	AccessKeySecret string  `yaml:"accessKeySecret"`
	Groups          []Group `yaml:"groups"`
}

var config Config

var defaultTestIPURL = "http://ifconfig.me"

//GetConfig 设置config
func GetConfig() *Config {
	return &config
}

func SetConfig(b *[]byte) error {
	err := yaml.UnmarshalStrict(*b, &config)
	if err != nil {
		return err
	}
	if config.TestIPURL == "" {
		config.TestIPURL = defaultTestIPURL
	}
	if config.WaitSeconds == 0 {
		config.WaitSeconds = 300
	}
	return nil
}
