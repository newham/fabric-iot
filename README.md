# fabric-iot 是一个基于Hyperledger Fabric框架实现的物联网访问控制管理的区块链模拟实验平台
[English](/README.en.md)  
> **WHAT WE HAVE**:  
> 1.采用ABAC作为权限控制方法，用golang实现了chaincode.  
> 2.编写了完整的脚本，轻松启动fabric-iot网络.  
> 3.详细的步骤说明，浅显易懂.    

## 0.准备工作
### 0.0.官网
[ hyperledger-fabric github ](https://github.com/hyperledger/fabric)  
[ hyperledger-fabric doc : 1.4 ](https://hyperledger-fabric.readthedocs.io/en/release-1.4/)  
[ node-sdk : 1.4 ](https://hyperledger.github.io/fabric-sdk-node/release-1.4/module-fabric-network.html)

### 0.1.操作系统
OS|版本
-|-
mac os| >= `10.13`
ubuntu| `16.04`  

### 0.2.软件：  
软件|版本
-|-
git|>= `2.0.0`
docker|>= `19`
docker-compose|>= `1.24`
node| `latest`
golang|>= `1.10`
*Hyperledger Fabric| `1.4.3`

***请务必保持 Hyperledger Fabric 版本一致, 更改版本，极易导致实验 `失败`，请谨慎选择！**

***请将golang、node的\bin加入PATH**

***本次实验基于Fabric 1.4.3，默认下载的Fabric版本是1.4.3**

***如果需要修改版本，请更改 `bootstrap.sh`**
```
VERSION=1.4.3

CA_VERSION=1.4.3
```
*****

### 0.3.下载本软件
```go
go get github.com/newham/fabric-iot
```

### 0.4.下载并配置Hyperledger Fabric
下载源码和bin  
```bash
./bootstrap.sh
```
export PATH
```bash
export PATH=$PATH:$(pwd)/fabric-samples/bin
```
***之后的诸如【启动网络、安装链码等操作需要先export PATH，这是让系统知道Hyperledger Fabric的bin目录】**  

***也可将其加入path中**

## 1.目录结构

|-chaincode....................................链码目录  
|--go....................................................golang编写的chaincode代码  
|-client............................................交互客户端  
|--nodejs..............................................nodejs编写的客户端代码  
|-network........................................fabric-iot网络配置和启动脚本  
|--channel-artifacts..............................channel配置文件，创世区块目录  
|--compose-files...................................docker-compose文件目录  
|--conn-conf.........................................链接docker网络配置目录  
|--crypto-config....................................存放证书、秘钥文件目录  
|--scripts...............................................cli运行脚本目录  
|--ccp-generate.sh...............................生成节点证书目录配置文件脚本  
|--ccp-template.json.............................配置文件json模板  
|--ccp-template.yaml.............................配置文件yaml模板  
|--clear-docker.sh................................清理停用docker镜像脚本  
|--configtx.yaml....................................configtxgen的配置文件  
|--crypto-config.yaml...........................生成证书、秘钥的配置文件  
|--down.sh............................................关闭网络shell脚本  
|--env.sh...............................................其他shell脚本import的公共配置  
|--generate.sh......................................初始化系统配置shell脚本  
|--install-cc.sh......................................安装链码shell脚本  
|--rm-cc.sh...........................................删除链码shell脚本  
|--up.sh................................................启动网络shell脚本  

# 2.分布式 `fabric-iot` 系统搭建步骤

## 2.1 虚拟机集群搭建

- 采用 `docker` 自带的 `swarm` 网络，让容器可夸主机相互访问

## 2.2 架构

host|ip|service
-|-|-
order|10.10.10.100|cli,order,org1.ca,org2.ca
peer1|10.10.10.101|peer0.org1,peer0.org2
peer2|10.10.10.102|peer1.org1,peer1.org2   

## 2.3 准备工作
软件|版本
-|-
[virtualbox](https://www.virtualbox.org/)|>=6.1
[vagrant](https://www.vagrantup.com/)|>=2.2.10

---
## 2.4 vagrant 建立虚拟机集群
进入工作目录
```
cd network/vms
```
启动虚拟机
```
./up.sh
```
进入虚拟机  
`order` 节点
```shell
ssh -p 2200 vagrant@10.10.10.100
```
`peer1` 节点
```shell
ssh -p 2201 vagrant@10.10.10.100
```
`peer2` 节点
```shell
ssh -p 2202 vagrant@10.10.10.100
```

## 3. docker swarm 建立分布式容器网络
### 3.1. 在 `order` 上执行，初始化一个管理员节点
`order` 节点
```shell
docker swarm init --advertise-addr 10.10.10.100 
```
接着生成其他节点加入集群的token, 注意：这里必需是 `manager` 角色才能共享网络
```shell
docker swarm join-token manager
```
输出:
```shell
To add a manager to this swarm, run the following command:

    docker swarm join --token xxxx 10.10.10.100:2377
```

### 2.2. 在 `peer1` 和 `peer2` 上执行，加入到 `swarm` 网络
`peer1`,`peer2` 节点
```shell
docker swarm join --token xxxx 10.10.10.100:2377
```
输出 `This node joined a swarm as a worker.`  
至此，三个node已经构建一个整体docker网络  
查看node  
```
docker node ls
```

### 3.3. 接下来在 `order` 上创建一个可跨主机访问的 `overlay` 网络 `iot_cross_net`
`order` 节点
```
docker network create --attachable --driver overlay iot_cross_net
```
*这个网络 `iot_cross_net` 属于 swarm 网络，可跨主机访问，在 `docker-compose` 文件中会用到。

`peer` 节点   
```
docker network list
```
*查看 `iot_cross_net` 是否创建成功

### 3.4. 启动服务
`order` 节点
```
./up-order.sh
```
`peer1` 节点
```
./up-peer1.sh
```
`peer2` 节点
```
./up-peer2.sh
```

### 3.5. 初始化channel
`order` 节点
```
./init-channel.sh 
```
成功输出：
```
========= All GOOD, [fabric-iot] execution completed =========== 


 _____   _   _   ____   
| ____| | \ | | |  _ \  
|  _|   |  \| | | | | | 
| |___  | |\  | | |_| | 
|_____| |_| \_| |____/  

```

### 3.6. 安装链码（chain code）
`order` 节点
```
./cc.sh install pc 1.0 go/pc Synchro
```

### 3.7. 测试（本地tool）
`order` 节点
```
./cc.sh invoke pc 1.0 go/pc AddPolicy '"{\"AS\":{\"userId\":\"13800010001\",\"role\":\"u1\",\"group\":\"g1\"},\"AO\":{\"deviceId\":\"D100010001\",\"MAC\":\"00:11:22:33:44:55\"}}"'
```
成功输出：`Chaincode invoke successful. result: status:200 payload:"OK"`

### 3.8. node客户端测试
*首先, 在宿主机上修改 `host` , 这一步非常重要，否则报错网络不通。（推荐使用 `SwitchHost!`）
```shell
# order
10.10.10.100 orderer.fabric-iot.edu
# peer1
10.10.10.101 peer0.org1.fabric-iot.edu
10.10.10.101 peer0.org2.fabric-iot.edu
# peer2
10.10.10.102 peer1.org2.fabric-iot.edu
10.10.10.102 peer1.org1.fabric-iot.edu
```
宿主机, `client/nodejs/`
```shell
# 初始化（生成用户）
./init.sh
```
测试调用链码
```
node invoke.js pc QueryPolicy 40db810e4ccb4cc1f3d5bc5803fb61e863cf05ea7fc2f63165599ef53adf5623
```


<br>
<br>
©2019 <a href="mailto:liuhanshmtu@163.com">liuhanshmtu@163.com</a> all rights reserved.
