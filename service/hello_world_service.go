package service

import (
	"fmt"
	"github.com/kfangw/embedded_contract/common/address"
)

var HelloWorldAddress = address.AddressFromString("HELLOWORLDADDRESS")

func Init() {
	Contracts[HelloWorldAddress] = RegisterAuthContract
}

func SetHelloWorld(es *EmbededService) ([]byte, error) {
	fmt.Println("SET HelloWorld")
	es.DB.Add("helloKey", "HelloWorld")
	return nil, nil
}

func GetHelloWorld(es *EmbededService) ([]byte, error) {
	fmt.Println("GET HelloWorld")
	result, _ := es.DB.Get("helloKey")
	return []byte(result), nil
}

func RegisterAuthContract(es *EmbededService) {
	es.Register("setHelloWorld", SetHelloWorld)
	es.Register("getHelloWorld", GetHelloWorld)
}
