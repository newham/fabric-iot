# Set default value
# Obtain the OS and Architecture string that will be used to select the correct
# native binaries for your platform, e.g., darwin-amd64 or linux-amd64
OS_ARCH=$(echo "$(uname -s | tr '[:upper:]' '[:lower:]' | sed 's/mingw64_nt.*/windows/')-$(uname -m | sed 's/x86_64/amd64/g')" | awk '{print tolower($0)}')
# timeout duration - the duration the CLI should wait for a response from
# another container before giving up
CLI_TIMEOUT=10
# default for delay between commands
CLI_DELAY=3
# compose files path
COMPOSE_FILE_PATH=compose-files
# use this as the default docker-compose yaml definition
COMPOSE_FILE=$COMPOSE_FILE_PATH/docker-compose-cli.yaml
#
COMPOSE_FILE_COUCH=$COMPOSE_FILE_PATH/docker-compose-couch.yaml
# org3 docker compose file
COMPOSE_FILE_ORG3=$COMPOSE_FILE_PATH/docker-compose-org3.yaml
# kafka and zookeeper compose file
COMPOSE_FILE_KAFKA=$COMPOSE_FILE_PATH/docker-compose-kafka.yaml
# two additional etcd/raft orderers
COMPOSE_FILE_RAFT2=$COMPOSE_FILE_PATH/docker-compose-etcdraft2.yaml
# certificate authorities compose file
COMPOSE_FILE_CA=$COMPOSE_FILE_PATH/docker-compose-ca.yaml
#
# use golang as the default language for chaincode
LANGUAGE=golang
# default image tag
IMAGETAG="latest"
# Versions of fabric known not to work with this release of first-network
BLACKLISTED_VERSIONS="^1\.0\. ^1\.1\.0-preview ^1\.1\.0-alpha"
# channel name
CHANNEL_NAME="iot-channel"
# system channem name
SYS_CHANNEL="iot-sys-channel"
# set exports
export COMPOSE_PROJECT_NAME=iot
export IMAGE_TAG=$IMAGE_TAG
export SYS_CHANNEL=$SYS_CHANNEL
# conn conf
CONN_CONF_PATH=conn-conf
# default consensus type,[kafka,solo,etcdraft]
CONSENSUS_TYPE="kafka"