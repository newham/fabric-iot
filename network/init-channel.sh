#!/bin/bash

# import env
. env.sh

# settings
CERTIFICATE_AUTHORITIES="true"
# not install chaincode
NO_CHAINCODE="true"
# use couchdb
IF_COUCHDB="couchdb"

function runScript() {
  # now run the end to end script
  docker exec cli scripts/script.sh $CHANNEL_NAME $CLI_DELAY $LANGUAGE $CLI_TIMEOUT $VERBOSE $NO_CHAINCODE
  if [ $? -ne 0 ]; then
    echo "ERROR !!!! Test failed"
    exit 1
  fi
}

runScript