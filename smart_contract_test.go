package tc

import (
	"tc/service"
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
	es.NativeCall(service.HelloWorldAddress, "getHelloWorld", []byte{})

}
