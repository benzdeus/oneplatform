package oneplatform

import (
	"github.com/benzdeus/oneplatform/configs"
	"github.com/benzdeus/oneplatform/env"
	_oneChatService "github.com/benzdeus/oneplatform/modules/onechat/services"
	_oneIDService "github.com/benzdeus/oneplatform/modules/oneid/services"
	one_entities "github.com/benzdeus/oneplatform/one_entities"
)

type Service struct {
	OneID   one_entities.OneIDService
	OneChat one_entities.OneChatService
}

type Options configs.Options
type OneChatRequestElementSendComponent one_entities.OneChatRequestElementSendComponent
type Choice one_entities.OneChatChoice

func Init(options Options) Service {
	env.Options = configs.Options(options)
	oneIDService := _oneIDService.NewOneIDService()
	oneChatService := _oneChatService.NewOneChatService()
	serviceResponse := Service{
		OneID:   oneIDService,
		OneChat: oneChatService,
	}

	return serviceResponse
}
