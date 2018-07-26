package service

import (
	"bytes"
	"errors"
	"fmt"
	"tc/common/address"
	"tc/database"
	"tc/operation"
)

type (
	Handler         func(*EmbededService) ([]byte, error)
	RegisterService func(*EmbededService)
)

var (
	Contracts = make(map[address.Address]RegisterService)
)

// Native service struct
// Invoke a native smart contract, new a native service
type EmbededService struct {
	DB         *database.DB
	ServiceMap map[string]Handler
	//Notifications []*event.NotifyEventInfo
	Code  []byte
	Input []byte
	//Tx            *types.Transaction
	Height uint32
	Time   uint32
	//ContextRef    context.ContextRef
}

func (es *EmbededService) Register(methodName string, handler Handler) {
	es.ServiceMap[methodName] = handler
}

func (es *EmbededService) Invoke() (interface{}, error) {
	bf := bytes.NewBuffer(es.Code)
	contract := new(operation.OperationInvoke)
	if err := contract.Deserialize(bf); err != nil {
		return false, err
	}
	services, ok := Contracts[contract.Address]
	if !ok {
		return false, fmt.Errorf("Native contract address %x haven't been registered.", contract.Address)
	}
	services(es)
	service, ok := es.ServiceMap[contract.Method]
	if !ok {
		return false, fmt.Errorf("Native contract %x doesn't support es function %s.",
			contract.Address, contract.Method)
	}
	args := es.Input
	es.Input = contract.Args
	//es.ContextRef.PushContext(&context.Context{ContractAddress: contract.Address})
	//notifications := es.Notifications
	//es.Notifications = []*event.NotifyEventInfo{}
	result, err := service(es)
	if err != nil {
		return result, errors.New("[Invoke] Native serivce function execute error!")
	}
	//es.ContextRef.PopContext()
	//es.ContextRef.PushNotifications(es.Notifications)
	//es.Notifications = notifications
	es.Input = args
	return result, nil
}

func (es *EmbededService) NativeCall(address address.Address, method string, args []byte) (interface{}, error) {
	bf := new(bytes.Buffer)
	c := operation.OperationInvoke{
		Address: address,
		Method:  method,
		Args:    args,
	}
	if err := c.Serialize(bf); err != nil {
		return nil, err
	}
	es.Code = bf.Bytes()
	return es.Invoke()
}
