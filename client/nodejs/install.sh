#!/bin/bash

# NODE_VERSION="12.15.0"

# echo "==> 切换到NODE_VERSION" 
# n use $NODE_VERSION

echo "==> 删除./node_modules" 
rm -rf ./node_modules

echo "==> 安装依赖" 
npm install