package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type SimpleChaincode struct {
}

//Patient

type Patient struct {
	Name       string `json:"name"`
	Id         string `json:"id"`
	Dob        string `json:"dob"`
	Bloodgroup string `json:"bloodgroup"`
}

//Doctor

type Doctor struct {
	Name       string `json:"name"`
	Id         string `json:"id"`
	Hospital   string `json:"hospital"`
	Department string `json:"department"`
}

//Pharmacy

type Pharmacy struct {
	Name  string `json:"name"`
	Id    string `json:"id"`
	Pin   string `json:"pin"`
	Owner string `json:"owner"`
}

//main

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}

// Init - initialize the chaincode

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("########### SimpleChain Init ###########")

	function, _ := stub.GetFunctionAndParameters()
	logger.info(function)
	if function != "init" {
		return shim.Error("Unknown function call")
	}

	return shim.Success(nil)
}

// Invoke - Our entry point for Invocations

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println(" ")
	fmt.Println("starting invoke, for - " + function)
	if function == "init" {
		return t.Init(stub)
	} else if function == "doc_create" {
		return t.doc_create(stub, args)
	} else if function == "doc_invoke" {
		return t.doc_invoke(stub, args)
	} else if function == "doc_query" {
		return t.doc_query(stub, args)
	} else if function == "doc_gethistory" {
		return t.doc_gethistory(stub, args)
	} else if function == "doc_querone" {
		return t.doc_queryone(stub, args)
	} else if function == "pat_create" {
		return t.pat_create(stub, args)
	} else if function == "pat_invoke" {
		return t.pat_invoke(stub, args)
	} else if function == "pat_query" {
		return t.pat_query(stub, args)
	} else if function == "pat_gethistory" {
		return t.pat_gethistory(stub, args)
	} else if function == "pat_queryone" {
		return t.pat_querone(stub, args)
	} else if function == "pharm_create" {
		return t.pharm_create(stub, args)
	} else if function == "pharm_invoke" {
		return t.pharm_invoke(stub, args)
	} else if function == "pharm_query" {
		return t.pharm_query(stub, args)
	} else if function == "pharm_gethistory" {
		return t.pharm_gethistory(stub, args)
	} else if function == "pharm_queryone" {
		return t.doc_create(stub, args)
	}

	return shim.Error("Function with the name " + function + " does not exist.")
}
