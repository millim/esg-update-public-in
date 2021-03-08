## 功能
阿里云上对服务器设置了白名单，但是部分IP不是固定IP，路由被重启后就不会被重新随机分配一个IP，因此弄了简单的内容对其进行更新。

## 注意
需要自己申请阿里云相关的权限


```bash
#-c 配置文件 -f 已有的ip地址文件 -p 进程pid文件
nohup esg -c config.yml -f public_ip >./log 2>&1 &

# config.yml
testUrl: 可选,测试公共IP的地址，http get 请求直接返回IP地址， 默认： http://ifconfig.me
updateUrl: 可选,当IP地址发生变化后，会调用的地址，http get 请求，会增加一个ip的query参数，待实现
waitSeconds: 可选,刷新周期，单位秒，默认:300
regionId: cn-hangzhou 阿里的区域
accessKeyId: 权限keyid
accessKeySecret: 权限keysecret
groups:
  - groupId: 安全组id
    port: 端口
    info: 描述
  - groupId: 安全组id
    port: 端口
    info: 描述
```