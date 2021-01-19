#!/bin/bash

# Exit on first error
set -e

if [ -n "$1" -a -n "$2" ]; then
    echo "$1$2" | sha256sum
else
    echo '[user_id] [device_id]' 
fi

