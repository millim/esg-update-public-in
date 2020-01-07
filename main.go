package main

import (
	"flag"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)


var file = "./public_ip"
var config Config

type Group struct {
	GroupId string `yaml:"groupId"`
	Port string `yaml:"port"`
	Info string `yaml:"info"`
}

type Config struct {
	RegionID string `yaml:"regionId"`
	AccessKeyID string `yaml:"accessKeyId"`
	AccessKeySecret string `yaml:"accessKeySecret"`
	Groups []Group `yaml:"groups"`

}

func main() {
	configAdd := ""
	flag.StringVar(&configAdd,"c","", "-c config.yml")
	flag.Parse()
	println(configAdd)

	if !fileExists(configAdd){
		println("no search config file, use -c config.yml")
		return
	}


	ymlFile, err := ioutil.ReadFile(configAdd)
	if err != nil {
		println("read config err--->",err)
		return
	}

	err = yaml.UnmarshalStrict(ymlFile, &config)
	if err != nil {
		println("config load error!")
		println(err)
		return
	}

	for {
		inspect()
		time.Sleep(5 * time.Minute)
	}

}

func inspect(){
	ip := getIP()
	if !fileExists(file) {
		createOriginalIP(ip)
	}else {
		oldIP, err := getOldIP()
		if err != nil {
			println(err.Error())
			return
		}
		if oldIP == ip{
			return
		}
		println(time.Now().String(),"----->","ip变了")
		updateAliyun(oldIP, ip)
		ioutil.WriteFile(file,[]byte(ip), 0644)
	}
}

func updateAliyun(oldIP, ip string){
	client, _ = ecs.NewClientWithAccessKey(config.RegionID, config.AccessKeyID, config.AccessKeySecret)
	if len(config.Groups) > 0 {
		for _, v := range config.Groups{
			updateGroup(v.GroupId, oldIP, ip, v.Port, v.Info)
		}
	}
}


func getOldIP()(string, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		println(err.Error())
		return "", err
	}
	return string(b), nil
}

func createOriginalIP(ip string){
	f, err := os.OpenFile(file, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		println(err.Error())
		return
	}
	defer f.Close()
	_, err = f.WriteString(ip)
	if err != nil {
		println(err.Error())
		return
	}
}


func fileExists(file string) bool{
	_, err := os.Stat(file)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}


func getIP() string{
	rep, err := http.Get("http://ifconfig.me")
	if err !=nil {
		return ""
	}

	defer rep.Body.Close()
	body, err := ioutil.ReadAll(rep.Body)
	if err !=nil {
		return ""
	}
	ip := string(body)
	return ip
}

var client *ecs.Client

func updateGroup(group, oldIP, ip, port, info string){
	removeGroup(group, oldIP, port)
	addGroup(group, ip, port, info)
}

func removeGroup(group, ip, port string){
	request := ecs.CreateRevokeSecurityGroupRequest()
	request.SecurityGroupId = group
	request.RegionId = config.RegionID
	request.IpProtocol = "tcp"
	request.PortRange = fmt.Sprintf("%s/%s",port,port)
	request.SourceCidrIp = ip
	_, err := client.RevokeSecurityGroup(request)
	if err != nil {
		fmt.Print(err.Error())
	}
}
func addGroup(group, ip, port, info string){
	request := ecs.CreateAuthorizeSecurityGroupRequest()
	request.SecurityGroupId = group
	request.RegionId = config.RegionID
	request.IpProtocol = "tcp"
	request.PortRange = fmt.Sprintf("%s/%s",port,port)
	request.SourceCidrIp = ip
	request.Description = info
	_, err := client.AuthorizeSecurityGroup(request)
	if err != nil {
		fmt.Print(err.Error())
	}

}
