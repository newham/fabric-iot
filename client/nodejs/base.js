//common
const fs = require('fs');
const path = require('path');
const ccpPath = path.resolve(__dirname, '..', '..', 'network', 'conn-conf', 'connection-org1.json');
const ccpJSON = fs.readFileSync(ccpPath, 'utf8');
const ccp = JSON.parse(ccpJSON);
//fabric SDK
const FabricCAServices = require('fabric-ca-client');
const { FileSystemWallet, Gateway, X509WalletMixin } = require('fabric-network');
//const value
const CHANNEL_NAME = 'iot-channel';
const POLICY_CC_NAME = 'policy-cc';
//exports
exports.ccp = ccp;
exports.ccpPath = ccpPath;
exports.path = path;
exports.FabricCAServices = FabricCAServices;
exports.FileSystemWallet = FileSystemWallet;
exports.X509WalletMixin = X509WalletMixin;
exports.Gateway = Gateway;
exports.CHANNEL_NAME = CHANNEL_NAME;
exports.POLICY_CC_NAME=POLICY_CC_NAME;
