package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
	"github.com/newham/fabric-iot/chaincode/go/m"
)

type AccessContract interface {
	Init(shim.ChaincodeStubInterface) sc.Response
	Invoke(shim.ChaincodeStubInterface) sc.Response
	CheckAccess(shim.ChaincodeStubInterface, []string) sc.Response
	Auth(string) m.ABACRequest
	GetAttrs(r m.ABACRequest) m.Attrs
	Synchro() sc.Response
}

type ChainCode struct {
	AccessContract
}

func NewAccessContract() AccessContract {
	return new(ChainCode)
}

func (cc *ChainCode) Auth(string) m.ABACRequest {
	return m.ABACRequest{}
}

func (cc *ChainCode) GetAttrs(r m.ABACRequest) m.Attrs {
	return m.Attrs{}
}

func (cc *ChainCode) CheckAccess(shim.ChaincodeStubInterface, []string) sc.Response {
	return shim.Success(m.OK)
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

func main() {
	// Create a new Smart Contract
	err := shim.Start(NewAccessContract())
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
