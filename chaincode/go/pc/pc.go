package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
	"github.com/newham/fabric-iot/chaincode/go/m"
)

type PolicyContract interface {
	Init(shim.ChaincodeStubInterface) sc.Response
	Invoke(shim.ChaincodeStubInterface) sc.Response
	Synchro() sc.Response
	CheckPolicy(m.Policy) bool
	AddPolicy(shim.ChaincodeStubInterface, []string) sc.Response
	UpdatePolicy(shim.ChaincodeStubInterface, []string) sc.Response
	QueryPolicy(shim.ChaincodeStubInterface, []string) sc.Response
	DeletePolicy(shim.ChaincodeStubInterface, []string) sc.Response
}

// Define the Smart Contract structure
type ChainCode struct {
}

/*
 * The Init method is called when the Smart Contract "fabcar" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (cc *ChainCode) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(m.OK)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "fabcar"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (cc *ChainCode) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "AddPolicy" {
		return cc.AddPolicy(APIstub, args)
	} else if function == "Synchro" {
		return cc.Synchro()
	} else if function == "QueryPolicy" {
		return cc.QueryPolicy(APIstub, args)
	} else if function == "DeletePolicy" {
		return cc.DeletePolicy(APIstub, args)
	} else if function == "UpdatePolicy" {
		return cc.UpdatePolicy(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func NewPolicyContract() PolicyContract {
	return new(ChainCode)
}

func (cc *ChainCode) parsePolicy(arg string) (m.Policy, error) {
	policyAsBytes := []byte(arg)
	policy := m.Policy{}
	err := json.Unmarshal(policyAsBytes, &policy)
	return policy, err
}

func (cc *ChainCode) CheckPolicy(p m.Policy) bool {
	return false
}

//This is the main smart contract of system
func (cc *ChainCode) AddPolicy(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of argumentcc. Expecting 1")
	}
	// parse policy
	policy, err := cc.parsePolicy(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	// check it
	if cc.CheckPolicy(policy) {
		return shim.Error("bad policy")
	}
	// put k-v to DB
	err = APIstub.PutState(policy.GetID(), policy.ToBytes())
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(m.OK)
}

func (cc *ChainCode) QueryPolicy(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of argumentcc. Expecting 1")
	}
	policyAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(policyAsBytes)
}

func (cc *ChainCode) UpdatePolicy(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of argumentcc. Expecting 1")
	}
	// parse policy
	policy, err := cc.parsePolicy(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	// check it
	if cc.CheckPolicy(policy) {
		return shim.Error("bad policy")
	}
	r := cc.QueryPolicy(APIstub, []string{policy.GetID()})
	if r.GetStatus() != 200 {
		return shim.Error("policy not exist")
	}
	return cc.AddPolicy(APIstub, args)
}

func (cc *ChainCode) DeletePolicy(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of argumentcc. Expecting 1")
	}
	err := APIstub.DelState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(m.OK)
}

func (cc *ChainCode) Synchro() sc.Response {
	return shim.Success(m.OK)
}

// The main function is only relevant in unit test mode. Only included here for completenescc.
func main() {

	// Create a new Smart Contract
	err := shim.Start(NewPolicyContract())
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
