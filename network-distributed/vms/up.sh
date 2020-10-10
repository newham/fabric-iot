#!/bin/bash
#
# Copyright Liu Han. All Rights Reserved.
#
# 基于vagrant的初始化脚本
#
# 失败退出
set -e
# 将用户的公钥加入到虚拟机
cp ~/.ssh/id_rsa.pub ./synced_folder/ssh
# 基于ubuntu/xenial64创建脚本
vagrant up