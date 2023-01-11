package services_test

import (
	"log"
	"os"
	"testing"

	"github.com/benzdeus/oneplatform/configs"
	"github.com/benzdeus/oneplatform/env"
	"github.com/benzdeus/oneplatform/modules/oneid/services"
)

func TestPostCheckOneMail(t *testing.T) {
	env.Options = configs.Options{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
	}

	service := services.NewOneIDService()

	t.Run("ShouldBeSuccess", func(t *testing.T) {

		response := service.PostCheckOneMail("ogstest555@one.th")

		if response.Result != "Success" {
			t.Error(response.Result)
			log.Println(response)
		}
	})

	t.Run("ShouldBeFail", func(t *testing.T) {
		response := service.PostCheckOneMail("asdasdas@one.th")

		if response.Result == "Success" {
			t.Error(response.Result)
		}
	})
}
