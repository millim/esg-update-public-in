## 功能
阿里云上对服务器设置了白名单，但是部分IP不是固定IP，路由被重启后就不会被重新随机分配一个IP，因此弄了简单的内容对其进行更新。


```bash
nohup esg -c config.yml >./log 2>&1 &

# config.yml
regionId: cn-hangzhou
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