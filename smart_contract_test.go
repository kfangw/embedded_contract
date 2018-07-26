package embedded_contract

import (
	"fmt"
	"github.com/kfangw/embedded_contract/service"
	"testing"
)

func TestSmartContract(t *testing.T) {

	config := &Config{
		Time:   10,
		Height: 10,
	}
	sc := SmartContract{
		Config: config,
	}

	es, _ := sc.NewEmbededService()

	service.Init()

	es.NativeCall(service.HelloWorldAddress, "setHelloWorld", []byte{})
	result, _ := es.NativeCall(service.HelloWorldAddress, "getHelloWorld", []byte{})

	fmt.Println(string(result.([]byte)))

}
