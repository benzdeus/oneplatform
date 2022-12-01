package services

import (
	"github.com/benzdeus/oneplatform/configs"
	"github.com/benzdeus/oneplatform/env"
)

type oneIDService struct {
	options configs.Options
}

func NewOneIDService() *oneIDService {
	return &oneIDService{
		options: env.Options,
	}
}
