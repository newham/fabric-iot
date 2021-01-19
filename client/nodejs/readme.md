# 客户端调用

## 0.配置HOSTS
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

## 1.安装node依赖
```shell
npm install
```

## 2.测试
宿主机
```shell
cd client/nodejs/
```
*进入客户端代码目录

```shell
# 初始化（生成用户）
./init.sh
```
测试调用链码
```
node invoke.js pc QueryPolicy 40db810e4ccb4cc1f3d5bc5803fb61e863cf05ea7fc2f63165599ef53adf5623
```