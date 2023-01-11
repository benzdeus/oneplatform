package services_test

import (
	"log"
	"os"
	"testing"

	"github.com/benzdeus/oneplatform/configs"
	"github.com/benzdeus/oneplatform/env"
	"github.com/benzdeus/oneplatform/modules/oneid/services"
)

func TestGetAccountDetailByIdNo(t *testing.T) {

	env.Options = configs.Options{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
	}
	service := services.NewOneIDService()
	data := service.LoginPWD(os.Getenv("USERNAME"), os.Getenv("PASSWORD"))

	t.Run("ShouldBeSuccess", func(t *testing.T) {

		response := service.GetAccountDetialByIdNo(os.Getenv("ID_NO"), data.AccessToken)

		if response.Result != "Success" {
			t.Errorf("%#+v", response.Data)
			log.Println(response)
		}
	})

	t.Run("ShouldBeFail", func(t *testing.T) {
		response := service.GetAccountDetialByIdNo("", data.AccessToken)

		if response.Result == "Success" {
			t.Error(response.Data)
		}
	})
}
