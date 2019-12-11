package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
	"github.com/newham/fabric-iot/chaincode/go/m"
)

type AccessContract interface {
	Init(shim.ChaincodeStubInterface) sc.Response
	Invoke(shim.ChaincodeStubInterface) sc.Response
	Synchro() sc.Response

	CheckAccess(shim.ChaincodeStubInterface, []string) sc.Response
	Auth(string) (m.ABACRequest, error)
}

type ChainCode struct {
	AccessContract
}

func NewAccessContract() AccessContract {
	return new(ChainCode)
}

func (cc *ChainCode) Auth(str string) (m.ABACRequest, error) {
	r := m.ABACRequest{}
	b := []byte(str)
	err := json.Unmarshal(b, &r)
	return r, err
}

func (cc *ChainCode) CheckAccess(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of argumentcc. Expecting 1")
	}
	r, err := cc.Auth(args[0])
	if err != nil {
		return shim.Error("403")
	}
	attrs := r.GetAttrs()
	//get policy
	resp := APIstub.InvokeChaincode("pc", [][]byte{[]byte("QueryPolicy"), []byte(attrs.GetId())}, "iot-channel")
	if resp.GetStatus() != 200 {
		return shim.Error("403")
	}

	policy := m.Policy{}
	err = json.Unmarshal(resp.GetPayload(), &policy)
	if err != nil {
		return shim.Error("403")
		// return shim.Error(attrs.GetId() + ";" + string(resp.GetPayload()) + ";" + err.Error())
	}
	//check AP
	if policy.AP != 1 {
		return shim.Error("403")
		// return shim.Error(string(policy.ToBytes()) + ": AP is deney")
	}
	//check AE
	if attrs.Timestamp > policy.AE.EndTime {
		//disable the contract
		// DeletePolicy(APIstub, attrs.GetId())
		return shim.Error("AE is timeout")
	}
	//get URL
	resp = APIstub.InvokeChaincode("dc", [][]byte{[]byte("GetURL"), []byte(attrs.DeviceId)}, "iot-channel")
	res, err := m.NewResource(resp.GetPayload())
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte(res.URL))
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
	if function == "CheckAccess" {
		return cc.CheckAccess(APIstub, args)
	} else if function == "Synchro" {
		return cc.Synchro()
	}

	return shim.Error("Invalid Smart Contract function name.")
}

// base function
func (cc *ChainCode) Synchro() sc.Response {
	return shim.Success(m.OK)
}

func main() {
	// Create a new Smart Contract
	err := shim.Start(NewAccessContract())
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
