/*
 * SPDX-License-Identifier: Apache-2.0
 */

'use strict';

const { ccpPath, path, FileSystemWallet, Gateway, CHANNEL_NAME } = require("./base")

function resp(status, msg) {
    return { status: status, msg: msg }
}

async function main() {
    try {
        const argv = process.argv;
        if (argv.length < 5) {
            // console.log("[chaincode name] [function name] [arg]");
            return;
        }
        // from 2 is args begin
        const ccName = argv[2]
        const fName = argv[3]
        const args = argv.slice(4) //[a,b] to "a b"

        invoke(ccName, fName, args)
    } catch (error) {
        // console.error(`Failed to evaluate transaction: ${error}`);
        process.exit(1);
    }
}
async function invoke(ccName, fName, args) {
    try {
        // Create a new file system based wallet for managing identities.
        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = new FileSystemWallet(walletPath);
        // console.log(`Wallet path: ${walletPath}`);

        // Check to see if we've already enrolled the user.
        const user = 'user1'
        const userExists = await wallet.exists(user);
        if (!userExists) {
            // console.log(`An identity for the user "${user}" does not exist in the wallet`);
            // console.log('Run the registerUser.js application before retrying');
            return resp(500, 'Run the registerUser.js application before retrying');
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccpPath, { wallet, identity: user, discovery: { enabled: true, asLocalhost: true } });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork(CHANNEL_NAME);

        // Get the contract from the network.
        const contract = network.getContract(ccName);

        // Evaluate the specified transaction.
        // queryCar transaction - requires 1 argument, ex: ('queryCar', 'CAR4')
        // queryAllCars transaction - requires no arguments, ex: ('queryAllCars')
        console.log(ccName, fName, ...args)
        switch (fName) {
            case "AddURL":
            case "AddPolicy":
            case "DeletePolicy":
            case "UpdatePolicy":
            case "CheckAccess":
                const r1 = await contract.submitTransaction(fName, ...args);
                console.log(`Transaction has been submit, result is: ${r1.toString()}`);
                return resp(200, r1.toString())
                break;
            default:
                const r2 = await contract.evaluateTransaction(fName, ...args);
                console.log(`Transaction has been evaluated, result is: ${r2.toString()}`);
                return resp(200, r2.toString())
                break;
        }

        // process.exit(1);
    } catch (error) {
        // console.error(`Failed to evaluate transaction: ${error}`);
        return resp(500, error)
        // process.exit(1);
    }
}

main();

module.exports = { invoke }
