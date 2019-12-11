#!/bin/bash

# Exit on first error
set -e

# don't rewrite paths for Windows Git Bash users
export MSYS_NO_PATHCONV=1
starttime=$(date +%s)
# use golang as programing language
CC_RUNTIME_LANGUAGE=golang
CC_SRC_PATH=github.com/newham/fabric-iot/chaincode/go/pc
CC_NAME=PC
CC_VERSION=1.0
CC_INVOKE_FUNC_NAME="Synchro"
# client container's name
CLI=cli
# action
ACTION=install

# set needed values
CONFIG_ROOT=/opt/gopath/src/github.com/hyperledger/fabric/peer
ORG1_MSPCONFIGPATH=${CONFIG_ROOT}/crypto/peerOrganizations/org1.fabric-iot.edu/users/Admin@org1.fabric-iot.edu/msp
ORG1_TLS_ROOTCERT_FILE=${CONFIG_ROOT}/crypto/peerOrganizations/org1.fabric-iot.edu/peers/peer0.org1.fabric-iot.edu/tls/ca.crt
ORG2_MSPCONFIGPATH=${CONFIG_ROOT}/crypto/peerOrganizations/org2.fabric-iot.edu/users/Admin@org2.fabric-iot.edu/msp
ORG2_TLS_ROOTCERT_FILE=${CONFIG_ROOT}/crypto/peerOrganizations/org2.fabric-iot.edu/peers/peer0.org2.fabric-iot.edu/tls/ca.crt
ORDERER_TLS_ROOTCERT_FILE=${CONFIG_ROOT}/crypto/ordererOrganizations/fabric-iot.edu/orderers/orderer.fabric-iot.edu/msp/tlscacerts/tlsca.fabric-iot.edu-cert.pem

if [ -n $1 -a -n $2 -a -n $3 -a -n $4 ]; then
  ACTION=$1
  CC_NAME=$2
  CC_VERSION=$3
  CC_SRC_PATH="github.com/newham/fabric-iot/chaincode/$4"
  if [ ! $5 = '' ]; then
    CC_INVOKE_FUNC_NAME=$5
  fi
  echo $ACTION ${CC_NAME}:${CC_VERSION} $CC_SRC_PATH $CC_INVOKE_FUNC_NAME
else
  echo './cc.sh ["install"/"upgrade"] [cc name] [cc version] [cc src path] [fname]'
  exit 1
fi

# 1:org 2:peer 3:port 4:ORG1_MSPCONFIGPATH 5:ORG1_TLS_ROOTCERT_FILE
function installCC() {
  if [ ${1} = 1 ]; then
    MSPCONFIGPATH=$ORG1_MSPCONFIGPATH
    TLS_ROOTCERT_FILE=$ORG1_TLS_ROOTCERT_FILE
  elif [ ${1} = 2 ]; then
    MSPCONFIGPATH=$ORG2_MSPCONFIGPATH
    TLS_ROOTCERT_FILE=$ORG2_TLS_ROOTCERT_FILE
  else
    echo "Worng Org,check it"
    exit 1
  fi
  echo "Installing smart contract on peer${2}.org${1}.fabric-iot.edu"
  docker exec \
    -e CORE_PEER_LOCALMSPID=Org${1}MSP \
    -e CORE_PEER_ADDRESS=peer${2}.org${1}.fabric-iot.edu:${3} \
    -e CORE_PEER_MSPCONFIGPATH=${MSPCONFIGPATH} \
    -e CORE_PEER_TLS_ROOTCERT_FILE=${TLS_ROOTCERT_FILE} \
    $CLI \
    peer chaincode install \
    -n "$CC_NAME" \
    -v "$CC_VERSION" \
    -p "$CC_SRC_PATH" \
    -l "$CC_RUNTIME_LANGUAGE"
}

function initCC() {
  echo "Instantiating smart contract on iot-channel"
  docker exec \
    -e CORE_PEER_LOCALMSPID=Org1MSP \
    -e CORE_PEER_MSPCONFIGPATH=${ORG1_MSPCONFIGPATH} \
    $CLI \
    peer chaincode instantiate \
    -o orderer.fabric-iot.edu:7050 \
    -C iot-channel \
    -n "$CC_NAME" \
    -l "$CC_RUNTIME_LANGUAGE" \
    -v "$CC_VERSION" \
    -c '{"Args":[]}' \
    -P "AND('Org1MSP.member','Org2MSP.member')" \
    --tls \
    --cafile ${ORDERER_TLS_ROOTCERT_FILE} \
    --peerAddresses peer0.org1.fabric-iot.edu:7051 \
    --tlsRootCertFiles ${ORG1_TLS_ROOTCERT_FILE}

  echo "Waiting for contract instantiating..."
  process 15
}

function process() {
  for ((i = 1; i <= ${1}; i++)); do
    sleep 1
    echo -e ">\c"
  done
  echo
}

function invokeCC() {
  echo "Submitting transaction:${CC_INVOKE_FUNC_NAME} to smart contract on iot-channel"
  echo "The transaction is sent to all of the peers so that chaincode is built before receiving the following requests"
  docker exec \
    -e CORE_PEER_LOCALMSPID=Org1MSP \
    -e CORE_PEER_MSPCONFIGPATH=${ORG1_MSPCONFIGPATH} \
    $CLI \
    peer chaincode invoke \
    -o orderer.fabric-iot.edu:7050 \
    -C iot-channel \
    -n "$CC_NAME" \
    -c '{"function":"'$CC_INVOKE_FUNC_NAME'","Args":[]}' \
    --waitForEvent \
    --tls \
    --cafile ${ORDERER_TLS_ROOTCERT_FILE} \
    --peerAddresses peer0.org1.fabric-iot.edu:7051 \
    --peerAddresses peer1.org1.fabric-iot.edu:8051 \
    --peerAddresses peer0.org2.fabric-iot.edu:9051 \
    --peerAddresses peer1.org2.fabric-iot.edu:10051 \
    --tlsRootCertFiles ${ORG1_TLS_ROOTCERT_FILE} \
    --tlsRootCertFiles ${ORG1_TLS_ROOTCERT_FILE} \
    --tlsRootCertFiles ${ORG2_TLS_ROOTCERT_FILE} \
    --tlsRootCertFiles ${ORG2_TLS_ROOTCERT_FILE}

}

function upgradeCC() {
  echo "Upgrading smart contract on iot-channel"
  docker exec \
    -e CORE_PEER_LOCALMSPID=Org1MSP \
    -e CORE_PEER_MSPCONFIGPATH=${ORG1_MSPCONFIGPATH} \
    $CLI \
    peer chaincode upgrade \
    -o orderer.fabric-iot.edu:7050 \
    -C iot-channel \
    -n "$CC_NAME" \
    -l "$CC_RUNTIME_LANGUAGE" \
    -v "$CC_VERSION" \
    -c '{"Args":[]}' \
    -P "AND('Org1MSP.member','Org2MSP.member')" \
    --tls \
    --cafile ${ORDERER_TLS_ROOTCERT_FILE} \
    --peerAddresses peer0.org1.fabric-iot.edu:7051 \
    --tlsRootCertFiles ${ORG1_TLS_ROOTCERT_FILE}

  echo "Waiting for instantiation request to be committed ..."
  sleep 15
}

# set -x

# install cc
installCC 1 0 7051
installCC 1 1 8051
installCC 2 0 9051
installCC 2 1 10051
# init
if [ $ACTION = 'install' ]; then
  initCC
elif [ $ACTION = 'upgrade' ]; then
  upgradeCC
fi
# try
if [ ! $CC_INVOKE_FUNC_NAME = '' ]; then
  invokeCC
fi

echo "done"

# set +x
