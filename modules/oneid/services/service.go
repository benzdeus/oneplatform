package services

import (
	"github.com/benzdeus/oneplatform/configs"
	"github.com/benzdeus/oneplatform/env"
)

type OneIDService struct {
	options configs.Options
}

func NewOneIDService() *OneIDService {
	return &OneIDService{
		options: env.Options,
	}
}
