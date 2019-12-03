package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

var OK = []byte("OK")

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the car structure, with 4 properties.  Structure tags are used by encoding/json library
type Policy struct {
	Object      string `json:"object"`
	Subject     string `json:"subject"`
	Permission  string `json:"permission"`
	Environment string `json:"environment"`
}

func (p *Policy) toString() string {
	b, err := json.Marshal(*p)
	if err != nil {
		return ""
	}
	return string(b)
}

/*
 * The Init method is called when the Smart Contract "fabcar" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(OK)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "fabcar"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "addPolicy" {
		return s.addPolicy(APIstub, args)
	} else if function == "initPolicy" {
		return s.initPolicy(APIstub)
	} else if function == "queryPolicy" {
		return s.queryPolicy(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

//This is the main smart contract of system
func (s *SmartContract) addPolicy(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	policyAsBytes := []byte(args[0])

	policy := Policy{}
	err := json.Unmarshal(policyAsBytes, &policy)

	if err != nil {
		return shim.Error(err.Error())
	}

	// put k-v to DB
	err = APIstub.PutState(policy.Object, policyAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(OK)
}

func (s *SmartContract) queryPolicy(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	policyAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(policyAsBytes)
}

func (s *SmartContract) initPolicy(APIstub shim.ChaincodeStubInterface) sc.Response {
	// p := Policy{"test_object", "test_subject", "allow", strconv.FormatInt(time.Now().Unix(), 10)}
	// s.addPolicy(APIstub, []string{p.toString()})
	s.addPolicy(APIstub, []string{`{"object":"test_object","subject":"test_subject","permission":"allow","environment":"1575377240"}`})
	return shim.Success(OK)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
