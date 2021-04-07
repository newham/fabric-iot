#!/bin/bash

# 更新软件
sudo apt-get update
# 安装依赖工具
sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common

# SET UP THE REPOSITORY
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

sudo add-apt-repository \
"deb [arch=amd64] https://download.docker.com/linux/ubuntu \
$(lsb_release -cs) \
stable"

# INSTALL DOCKER ENGINE
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io docker-compose

#将当前用户添加至docker用户组
sudo gpasswd -a $USER docker
#更新docker用户组
newgrp docker

#添加淘宝源
echo "==> 添加淘宝源"
# sudo echo '{ "registry-mirrors": ["https://f1z25q5p.mirror.aliyuncs.com"] }' > /etc/docker/daemon.json
sudo cp daemon.json /etc/docker/
echo "==> 重启docker以适配"
sudo systemctl daemon-reload
sudo systemctl restart docker