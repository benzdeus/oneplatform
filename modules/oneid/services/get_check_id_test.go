package services_test

import (
	"log"
	"testing"

	"github.com/benzdeus/oneplatform/modules/oneid/services"
)

func TestGetCheckId(t *testing.T) {
	service := services.NewOneIDService()

	t.Run("ShouldBeSuccess", func(t *testing.T) {

		response := service.GetCheckId("")

		if response.Message != "Success" {
			t.Error(response.Message)
			log.Println(response)
		}
	})

	t.Run("ShouldBeFail", func(t *testing.T) {
		response := service.GetCheckId("")

		if response.Message == "Success" {
			t.Error(response.Message)
		}
	})
}
