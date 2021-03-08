package action

import (
	"ali-ecs-public-ip/lib/config"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"os"
)

func CreateOriginalIP(ip string) {
	f, err := os.OpenFile(config.PublicIPPath(), os.O_CREATE|os.O_RDWR, 0644)
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

var client *ecs.Client

func UpdateAliIP(oldIP, ip string) {
	config := config.GetConfig()
	client, _ = ecs.NewClientWithAccessKey(config.RegionID, config.AccessKeyID, config.AccessKeySecret)
	if len(config.Groups) > 0 {
		for _, v := range config.Groups {
			updateGroup(v.GroupId, oldIP, ip, v.Port, v.Info)
		}
	}
}

func updateGroup(group, oldIP, ip, port, info string) {
	removeGroup(group, oldIP, port)
	addGroup(group, ip, port, info)
}

func removeGroup(group, ip, port string) {
	if ip == "" {
		return
	}
	request := ecs.CreateRevokeSecurityGroupRequest()
	request.SecurityGroupId = group
	request.RegionId = config.GetConfig().RegionID
	request.IpProtocol = "tcp"
	request.PortRange = fmt.Sprintf("%s/%s", port, port)
	request.SourceCidrIp = ip
	_, err := client.RevokeSecurityGroup(request)
	if err != nil {
		fmt.Print(err.Error())
	}
}

func addGroup(group, ip, port, info string) {
	if ip == "" {
		return
	}
	request := ecs.CreateAuthorizeSecurityGroupRequest()
	request.SecurityGroupId = group
	request.RegionId = config.GetConfig().RegionID
	request.IpProtocol = "tcp"
	request.PortRange = fmt.Sprintf("%s/%s", port, port)
	request.SourceCidrIp = ip
	request.Description = info
	_, err := client.AuthorizeSecurityGroup(request)
	if err != nil {
		fmt.Print(err.Error())
	}

}
