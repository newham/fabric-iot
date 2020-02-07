# fabric-iot is a Blockchain Based Decentralized Access Control System in IoT 
[中文版](/README.md)  
> **WHAT WE HAVE**:  
> 1.Using ABAC as permission control method，and chaincode is writed with golang.  
> 2.The Hyperledger Fabric network is easily started with a complete shell script.  
> 3.Detailed instructions, easy to understand.    

## 0.Environment & Preparation
### 0.1.Operation System
OS|VETSION
-|-
mac os|>=10.13
ubuntu|16.04  

### 0.2.Software：  
SDK|VERSION
-|-
git|>=2.0.0
docker|>=19
docker-compose|>=1.24
node|>=12
golang|>=1.10
Hyperledger Fabric|>=1.4

### 0.3.Download src
```go
go get github.com/newham/fabric-iot
```

## 1.Directory Structure

|-chaincode....................................chaincode folder  
|--go..........................................chaincode in golang  
|-client.......................................client of blockchain  
|--nodejs......................................code in nodejs of client  
|-network......................................fabric-iot network module folder,genesis.block  
|--channel-artifacts...........................channel comfig folder  
|--compose-files...............................docker-compose folder  
|--conn-conf...................................connection config of docker  
|--crypto-config...............................certs folder    
|--scripts.....................................shell scrpits of cli  
|--ccp-generate.sh.............................generate certs and blockchain config shell  
|--ccp-template.json...........................json template of connection config  
|--ccp-template.yaml...........................template of connection config  
|--clear-docker.sh.............................clearn docker images  
|--configtx.yaml...............................config of configtxgen  
|--crypto-config.yaml..........................config of ccp-generate.sh  
|--down.sh.....................................shutdown the network  
|--env.sh......................................environment of other shell scripts  
|--generate.sh.................................init config  
|--install-cc.sh...............................install chaincode  
|--rm-cc.sh....................................delete chaincode  
|--up.sh.......................................start network  
## 2.Build Network of Fabric
### 2.1.Generate cers and configs
open folder  
```shell
cd network
```
run  
```shell
./generate.sh
```
cers are saved to `crypto-config`
### 2.2.Start Network
```shell
./up.sh
```
### 2.3.Install Chaincode
```shell
./cc-install.sh
```
### 2.4.Upgrade Chaincode
```shell
./cc-upgrade.sh [new_version]
```
*chaincode is saved to `/chaincode/go`，we only support write code with golang now.
### 2.5.Close Network
```shell
./down.sh
```
*all the stoped docker containers or images will be delete ，be carefull before you do this

## 3.Interacting with the Blockchain
### 3.1.Init Code
enter the foder   
*we have only a client writed by nodejs  
```shell
cd client/nodejs
```
install modules of this program
```shell
npm install
```
### 3.2.Create Users
create admin  
```shell
node ./enrollAdmin.js
```
create user by admin
```shell
node ./registerUser.js
```
### 3.3.Run Chaincode
```shell
node ./invoke.js [chaincode_name] [function_name] [args]
```

<br>
<br>
©2019 liuhan[liuhanshmtu@163.com] all rights reserved.
