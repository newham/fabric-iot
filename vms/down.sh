#!/bin/bash
#
# Copyright Liu Han. All Rights Reserved.
#
# 基于vagrant的初始化脚本
#

# 基于ubuntu/xenial64创建脚本
if [ "$1" = '-ff' ];then
    echo '==> 虚拟机:清除'
    vagrant destroy -f
    echo '==> 缓存:删除'
    rm -rf .vagrant
elif [ "$1" = '-f' ];then
    echo '==> 虚拟机:关机'
    vagrant halt
else
    echo '==> 虚拟机:挂起'
    vagrant suspend
fi