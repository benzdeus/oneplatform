package services

import (
	"testing"

	"github.com/benzdeus/oneplatform/configs"
	"github.com/benzdeus/oneplatform/env"
)

func TestOneIDService_GetAccountData(t *testing.T) {
	env.Options = configs.Options{
		ClientID:     0,
		ClientSecret: "",
	}

	token := ``

	oneidService := NewOneIDService()
	res := oneidService.GetAccountData(token)

	t.Errorf("%#v", res)

}
