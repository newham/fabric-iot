#!/bin/bash
#
# Copyright Liu Han. All Rights Reserved.
#
# 基于vagrant的初始化脚本
#
# 失败退出
set -e

SDK_DIR="./sdk/"

# 检查用户是否有秘钥对
if [ ! -d "~/.ssh/id_rsa.pub" ];then
    echo "请先创建ssh秘钥对，[ssh-keygen]"
    exit 1
fi

# 如果没有公钥则复制
if [ ! -d "${SDK_DIR}id_rsa.pub" ];then
    # 如果.ssh没有公钥就创建一个
    cp ~/.ssh/id_rsa.pub ${SDK_DIR}
fi
# 基于ubuntu/xenial64创建脚本
echo '==> 启动虚拟机'
vagrant up