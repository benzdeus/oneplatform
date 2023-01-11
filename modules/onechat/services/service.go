package services

import (
	"github.com/benzdeus/oneplatform/configs"
	"github.com/benzdeus/oneplatform/env"
)

type OneChatService struct {
	options configs.Options
}

func NewOneChatService() *OneChatService {
	return &OneChatService{
		options: env.Options,
	}
}
