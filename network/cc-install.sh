#!/bin/bash

# Exit on first error
set -e

./cc.sh install ac 1.0 go/ac Synchro
./cc.sh install pc 1.0 go/pc Synchro
./cc.sh install dc 1.0 go/dc Synchro

