package services

import (
	"github.com/benzdeus/oneplatform/configs"
	"github.com/benzdeus/oneplatform/env"
	"testing"
)

func TestOneIDService_GetAccountData(t *testing.T) {
	env.Options = configs.Options{
		ClientID:     0,
		ClientSecret: "xxx",
	}
}
