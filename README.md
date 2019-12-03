# fabric-iot 是一个基于Hyperledger Fabric框架实现的物联网访问控制管理的区块链模拟实验平台
[English](/README.en.md)  
> **WHAT WE HAVE**:  
> 1.采用ABAC作为权限控制方法，用golang实现了chaincode.  
> 2.编写了完整地脚本，轻松启动Hyperledger Fabric网络.  
> 3.详细的步骤说明，浅显易懂.    

## 0.准备工作
### 0.1.操作系统
OS|版本
-|-
mac os|>=10.13
ubuntu|16.04  

### 0.2.软件：  
软件|版本
-|-
git|>=2.0.0
docker|>=19
docker-compose|>=1.24
node|>=12
golang|>=1.10

## 1.目录结构

|-chaincode....................................链码目录  
|--go..........................................golang编写的chaincode代码  
|-client.......................................交互客户端  
|--nodejs......................................nodejs编写的客户端代码  
|-network......................................fabric-iot网络配置和启动脚本  
|--channel-artifacts...........................channel配置文件，创世区块目录  
|--compose-files...............................docker-compose文件目录  
|--conn-conf...................................链接docker网络配置目录  
|--crypto-config...............................存放证书、秘钥文件目录  
|--scripts.....................................cli运行脚本目录  
|--ccp-generate.sh.............................生成节点证书目录配置文件脚本  
|--ccp-template.json...........................配置文件json模板  
|--ccp-template.yaml...........................配置文件yaml模板  
|--clear-docker.sh.............................清理停用docker镜像脚本  
|--configtx.yaml...............................configtxgen的配置文件  
|--crypto-config.yaml..........................生成证书、秘钥的配置文件  
|--down.sh.....................................关闭网络shell脚本  
|--env.sh......................................其他shell脚本import的公共配置  
|--generate.sh.................................初始化系统配置shell脚本  
|--install-cc.sh...............................安装链码shell脚本  
|--rm-cc.sh....................................删除链码shell脚本  
|--up.sh.......................................启动网络shell脚本  
## 2.搭建网络
### 2.1.生成证书、配置文件等
进入目录  
```shell
cd network
```
运行脚本  
```shell
./generate.sh
```
证书保存在`crypto-config`中
### 2.2.启动网络
```shell
./up.sh
```
### 2.3.安装链码
```shell
./install-cc.sh
```
*chaincode被保存在`/chaincode/go`目录中，目前只用golang实现
### 2.4.关闭网络
```shell
./down.sh
```
*每次关闭网络会删除所有docker容器和镜像，请谨慎操作

## 3.与区块链交互
### 3.1.初始化代码
进入客户端代码目录  
*目前只实现了nodejs的客户端  
```shell
cd client/nodejs
```
安装依赖
```shell
npm install
```
### 3.2.创建用户
创建管理员账户
```shell
node ./enrollAdmin.js
```
用管理员账户创建子用户
```shell
node ./registerUser.js
```
### 3.3.调用chaincode
```shell
node ./invoke.js [chaincode_name] [function_name]
```

<br>
<br>
©2019 liuhan[liuhanshmtu@163.com] all rights reserved.
