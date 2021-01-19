#!/bin/bash

# 默认筛选关键词
key=iot
# 输入关键词
if [ -n "$1" ] ;then
    key="$1" 
fi

# 取得所有符合条件的容器名
containers=$(docker ps -a  |grep $key |awk '{print $1}')
if [ ! -n "$containers" ] ;then
    echo "no containers"
else
    # 删除停止的容器
    docker rm $containers
fi

# 取得所有符合条件的镜像名
imgs=$(docker images |grep $key |awk '{print $3}')
if [ ! -n "$imgs" ] ;then
    echo "no imgs"
else
    # 删除镜像
    docker rmi -f $imgs
fi
