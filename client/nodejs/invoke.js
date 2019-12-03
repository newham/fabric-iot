/*
 * SPDX-License-Identifier: Apache-2.0
 */

'use strict';

const { ccpPath, path, FileSystemWallet, Gateway, CHANNEL_NAME, POLICY_CC_NAME } = require("./base")

async function main() {
    try {
        const argv = process.argv;
        if (argv.length < 3) {
            console.log("Input a function");
            return;
        }
        // from 2 is args begin
        const fName = argv[2]
        switch (fName) {
            case "queryPolicy":
                if (argv.length < 4) {
                    console.log("Query need a key")
                    return
                }
                const key = argv[3]
                queryPolicy(fName, key)
                break;
            default:
                console.log("invoke support : queryPolicy, ")
                break
        }
    } catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`);
        process.exit(1);
    }
}
async function queryPolicy(fName, key) {
    try {
        // Create a new file system based wallet for managing identities.
        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = new FileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        // Check to see if we've already enrolled the user.
        const userExists = await wallet.exists('user1');
        if (!userExists) {
            console.log('An identity for the user "user1" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return;
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccpPath, { wallet, identity: 'user1', discovery: { enabled: true, asLocalhost: true } });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork(CHANNEL_NAME);

        // Get the contract from the network.
        const contract = network.getContract(POLICY_CC_NAME);

        // Evaluate the specified transaction.
        // queryCar transaction - requires 1 argument, ex: ('queryCar', 'CAR4')
        // queryAllCars transaction - requires no arguments, ex: ('queryAllCars')
        console.log(fName, key)
        const result = await contract.evaluateTransaction(fName, key);
        console.log(`Transaction has been evaluated, result is: ${result.toString()}`);
    } catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`);
        process.exit(1);
    }
}

main();
