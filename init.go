package oneplatform

import (
	"github.com/benzdeus/oneplatform/configs"
	"github.com/benzdeus/oneplatform/env"
	"github.com/benzdeus/oneplatform/modules/oneid/services"
)

type service struct {
	OneID *services.OneIDService
}

type Options configs.Options

func Init(options interface{}) service {
	env.Options = options.(configs.Options)
	oneIDService := services.NewOneIDService()
	serviceResponse := service{
		OneID: oneIDService,
	}

	return serviceResponse
}
