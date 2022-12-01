package services_test

import (
	"testing"

	"github.com/benzdeus/oneplatform/configs"
	"github.com/benzdeus/oneplatform/env"
	"github.com/benzdeus/oneplatform/modules/oneid/services"
)

func TestLoginPWDServiceSuccess(t *testing.T) {
	env.Options = configs.Options{
		ClientID:     0,
		ClientSecret: "xxx",
	}

	var request struct {
		UserName string
		Password string
	}

	request.UserName = "xxx"
	request.Password = "xxx"

	t.Run("Login Should be success", func(t *testing.T) {

		service := services.NewOneIDService()
		response := service.LoginPWD(request.UserName, request.Password)

		if response.ResponseCode != 200 {
			t.Errorf("%#v", response)
		}
	})

	t.Run("Login Should be fail", func(t *testing.T) {

		service := services.NewOneIDService()
		response := service.LoginPWD(request.UserName, request.Password)

		if response.ResponseCode == 200 {
			t.Errorf("%#v", response)
		}

	})
}
