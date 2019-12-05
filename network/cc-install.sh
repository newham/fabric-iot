#!/bin/bash

# Exit on first error
set -e

set -x
./cc.sh install $1 $2 $3 $4
set +x