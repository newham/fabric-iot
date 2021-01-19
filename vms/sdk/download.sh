

# 下载go和hyperledger bin (go 不是必须的)
# if [ -d "go1.10.8.linux-amd64.tar.gz" ]; then
#     wget https://studygolang.com/dl/golang/go1.10.8.linux-amd64.tar.gz
# fi
if [ -d "hyperledger-fabric-linux-amd64-1.4.3.tar.gz" ]; then
    wget https://github.com/hyperledger/fabric/releases/download/v1.4.3/hyperledger-fabric-linux-amd64-1.4.3.tar.gz
fi

# mkdir go 
# tar xzvf go1.10.8.linux-amd64.tar.gz -C ./go

mkdir hyperledger-fabric
tar zxvf hyperledger-fabric-linux-amd64-1.4.3.tar.gz -C ./hyperledger-fabric

#!/bin/bash
echo "Set Fabric binaries [$(pwd)/hyperledger-fabric/bin] to PATH"
# export PATH=$PATH:$(pwd)/hyperledger-fabric/bin
echo "export PATH=\$PATH:$(pwd)/hyperledger-fabric/bin" >> ~/.bashrc