> 请确保本机用户已经生成私钥对，`ssh-keygen` 命令可生成

> 公钥以 `id_rsa.pub` 的形式被拷贝到本目录，用于无密码登录vms

> 其他公共脚本在 `sdk` 下

1.安装软件和环境
```shell
./bootstrap.sh
```
*将自动执行:   
- docker 安装
- docker 镜像拉取
- hyperledger bin 安装
- export PATH