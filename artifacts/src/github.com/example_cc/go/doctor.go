package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type SimpleChaincode struct {
}

// query
// Every readonly functions in the ledger will be here
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### example_cc0 Init ###########")

	_, args := stub.GetFunctionAndParameters()
	logger.info("=========In doc_create========")
	// var newDoc Doc
	// json.Unmarshal([]byte(args[2]), &newDoc)
	// var doc = Doc{Name: newDoc.Name, Id: newDoc.Id, Hospital: newDoc.Hospital, Department: newDoc.Department}
	// docAsBytes, _ := json.Marshal(doc)
	stub.PutState(args[0], []byte(args[1]))

	// Notify listeners that an event "eventInvoke" have been executed (check line 19 in the file invoke.go)
	err := stub.SetEvent("eventCreateDoc", []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)

}
func (t *SimpleChaincode) doc_query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("########### Doctor query ###########")

	// Check whether the number of arguments is sufficient
	if len(args) < 2 {
		return shim.Error("The number of arguments is insufficient.")
	}

	// Like the Invoke function, we manage multiple type of query requests with the second argument.

	if args[1] == "all" {

		// GetState by passing lower and upper limits
		resultsIterator, err := stub.GetStateByRange("", "")
		if err != nil {
			return shim.Error(err.Error())
		}
		defer resultsIterator.Close()

		// buffer is a JSON array containing QueryResults
		var buffer bytes.Buffer
		buffer.WriteString("[")

		bArrayMemberAlreadyWritten := false
		for resultsIterator.HasNext() {
			queryResponse, err := resultsIterator.Next()
			if err != nil {
				return shim.Error(err.Error())
			}

			if bArrayMemberAlreadyWritten == true {
				buffer.WriteString(",")
			}
			buffer.WriteString("{\"Key\":")
			buffer.WriteString("\"")
			buffer.WriteString(queryResponse.Key)
			buffer.WriteString("\"")

			buffer.WriteString(", \"Record\":")
			// Record is a JSON object, so we write as-is
			buffer.WriteString(string(queryResponse.Value))
			buffer.WriteString("}")
			bArrayMemberAlreadyWritten = true
		}

		buffer.WriteString("]")

		fmt.Printf("- queryAllDoc:\n%s\n", buffer.String())

		return shim.Success(buffer.Bytes())
	}

	// If the arguments given don’t match any function, we return an error

	return shim.Error("Unknown query action, check the second argument.")
}

// invoke
// Every functions that read and write in the ledger will be here

func (t *SimpleChaincode) doc_invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("########### Doctor invoke ###########")

	if len(args) < 2 {
		return shim.Error("The number of arguments is insufficient.")
	}

	// Changing details of Doctor by Accepting Key and Value

	if args[1] == "changeDoc" && len(args) == 4 {

		docAsBytes, _ := stub.GetState(args[2])
		doc := Doc{}

		json.Unmarshal(docAsBytes, &doc)
		doc.Owner = args[3]

		docAsBytes, _ = json.Marshal(doc)
		stub.PutState(args[2], docAsBytes)

		// Notify listeners that an event "eventInvoke" have been executed (check line 19 in the file invoke.go)
		err := stub.SetEvent("eventChangeDoc", []byte{})
		if err != nil {
			return shim.Error(err.Error())
		}

		return shim.Success(nil)
	}

	// Updating all fields of record

	if args[1] == "updateRecord" && len(args) == 4 {
		fmt.Println("Update All")
		var newDoc Doc
		json.Unmarshal([]byte(args[3]), &newDoc)
		var doc = Doc{Name: newDoc.Name, Id: newDoc.Id, Quality: newDoc.Quality, Owner: newDoc.Owner}
		docAsBytes, _ := json.Marshal(doc)

		// Updating Record

		stub.PutState(args[2], docAsBytes)

		// Notify listeners that an event "eventInvoke" have been executed (check line 19 in the file invoke.go)

		err := stub.SetEvent("eventUpdateRecords", []byte{})
		if err != nil {
			return shim.Error(err.Error())
		}

		return shim.Success(nil)
	}

	// If the arguments given don’t match any function, we return an error

	return shim.Error("Unknown invoke action, check the second argument.")
}

//  Retrieves a single record from the ledger by accepting Key value

func (t *SimpleChaincode) doc_queryone(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 2 {
		return shim.Error("The number of arguments is insufficient.")
	}

	// GetState retrieves the data from ledger using the Key

	docAsBytes, _ := stub.GetState(args[1])

	// Transaction Response

	return shim.Success(docAsBytes)

}

// Adds a new transaction to the ledger

func (s *SimpleChaincode) doc_create(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	logger.info("=========In doc_create========")
	var newDoc Doc
	json.Unmarshal([]byte(args[2]), &newDoc)
	var doc = Doc{Name: newDoc.Name, Id: newDoc.Id, Hospital: newDoc.Hospital, Department: newDoc.Department}
	docAsBytes, _ := json.Marshal(doc)
	stub.PutState(args[1], docAsBytes)

	// Notify listeners that an event "eventInvoke" have been executed (check line 19 in the file invoke.go)
	err := stub.SetEvent("eventCreateDoc", []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// Get History of a transaction by passing Key

func (s *SimpleChaincode) doc_gethistory(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	docKey := args[1]
	fmt.Printf("##### start History of Record: %s\n", docKey)

	resultsIterator, err := stub.GetHistoryForKey(docKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the details of doctor
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON )
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getHistoryForDoctor returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}
