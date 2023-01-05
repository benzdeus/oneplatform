package services_test

import (
	"os"
	"testing"

	"github.com/benzdeus/oneplatform/configs"
	"github.com/benzdeus/oneplatform/env"
	"github.com/benzdeus/oneplatform/modules/oneid/services"
)

func TestPostDeleteEmail(t *testing.T) {

	t.Run("ShouldBeSuccess", func(t *testing.T) {

		env.Options = configs.Options{
			ClientID:     os.Getenv("CLIENT_ID"),
			ClientSecret: os.Getenv("CLIENT_SECRET"),
		}

		service := services.NewOneIDService()
		data := service.LoginPWD("", "")

		response := service.PostDeleteEmail("@gmail.com", data.AccessToken)

		if response.Result != "Success" {
			t.Error(response.Result)
		}
	})
	// t.Run("ShouldBeFail", func(t *testing.T) {
	// 	response := service.PostDeleteEmail("@gmail.com", data.AccessToken)

	// 	if response.Data == "Success" {
	// 		t.Error(response.Result)
	// 	}
	// })
}
