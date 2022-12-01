package oneplatform

import (
	"github.com/benzdeus/oneplatform/configs"
	"github.com/benzdeus/oneplatform/env"
	"github.com/benzdeus/oneplatform/modules/oneid/services"
)

type service struct {
	OneID *services.OneIDService
}

func Init(options configs.Options) service {
	env.Options = options
	oneIDService := services.NewOneIDService()
	serviceResponse := service{
		OneID: oneIDService,
	}

	return serviceResponse
}
