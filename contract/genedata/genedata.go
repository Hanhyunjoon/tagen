package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}

type UserRating struct{
	 Skin `json:"skin"`
}
type Skin struct{
	Aging int `json:"aging"`
	Elasticity int `json:"elasticity"`
	Pigmentation int `json:"Pigmentation"`
}
// type Obesity struct{
	
// }
// type Baldness struct{
	
// }
// type Healthcare struct{
	
// }


func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	function, args := APIstub.GetFunctionAndParameters()

	 if function == "addGene" {
		return s.addGene(APIstub, args)
	} else if function == "getGene" {
		return s.getGene(APIstub, args)
	} 
	return shim.Error("Invalid Smart Contract function name.")
}
 
func (s *SmartContract) addGene(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 4 {
		return shim.Error("fail!")
	}
	aging,_:=strconv.Atoi(args[1])
	 elasticity,_:=strconv.Atoi(args[2])
	 pigmentation,_:=strconv.Atoi(args[3])

	
	var user = UserRating{Skin{aging,elasticity,pigmentation}}
	userAsBytes, _ := json.Marshal(user)
	APIstub.PutState(args[0], userAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) getGene(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	UserAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(UserAsBytes)
}

func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}