package services_test

import (
	"testing"

	"github.com/benzdeus/oneplatform/configs"
	"github.com/benzdeus/oneplatform/env"
	"github.com/benzdeus/oneplatform/modules/oneid/services"
)

func TestPostCheckUsernameAndMobile(t *testing.T) {

	env.Options = configs.Options{
		ClientID:     "",
		ClientSecret: "",
	}

	service := services.NewOneIDService()
	t.Run("ShouldBeSuccess", func(t *testing.T) {
		response := service.PostCheckUsernameAndMobile("ogstest555", "0999999999")

		if response.Data == false {
			t.Error(response.Result)
		}
	})
	t.Run("ShouldBeFail", func(t *testing.T) {
		response := service.PostCheckUsernameAndMobile("ogstest555", "0911111111")

		if response.Data == true {
			t.Error(response.Result)
		}
	})
}
