package services_test

import (
	"os"
	"testing"

	"github.com/benzdeus/oneplatform/configs"
	"github.com/benzdeus/oneplatform/env"
	"github.com/benzdeus/oneplatform/modules/oneid/services"
)

func TestPostAddAddress(t *testing.T) {

	env.Options = configs.Options{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
	}

	service := services.NewOneIDService()
	data := service.LoginPWD(os.Getenv("USERNAME"), os.Getenv("PASSWORD"))

	t.Run("ShouldBeSuccess", func(t *testing.T) {

		response := service.PostAddAddress("1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1", "Y", data.AccessToken)

		if response.Result != "Success" {
			t.Errorf("%#+v", response)
		}
	})
	t.Run("ShouldBeFail", func(t *testing.T) {
		response := service.PostAddAddress("1", "1", "1", "1", "1", "1", "", "1", "1", "1", "1", "1", "1", "", data.AccessToken)

		if response.Data == "Success" {
			t.Error(response.Result)
		}
	})
}
