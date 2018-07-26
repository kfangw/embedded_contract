package embedded_contract

import (
	"github.com/kfangw/embedded_contract/database"
	"github.com/kfangw/embedded_contract/service"
)

type SmartContract struct {
	Config *Config
}

type Config struct {
	Time   uint32
	Height uint32
}

func (sc *SmartContract) NewEmbededService() (*service.EmbededService, error) {
	service := &service.EmbededService{
		DB:         database.NewDB(),
		Time:       sc.Config.Time,
		Height:     sc.Config.Height,
		ServiceMap: make(map[string]service.Handler),
	}
	return service, nil
}
