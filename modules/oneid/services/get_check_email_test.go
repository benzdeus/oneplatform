package services_test

import (
	"log"
	"testing"

	"github.com/benzdeus/oneplatform/modules/oneid/services"
)

func TestGetCheckEmail(t *testing.T) {
	// env.Options = configs.Options{
	// 	ClientID:     os.Getenv("CLIENT_ID"),
	// 	ClientSecret: os.Getenv("CLIENT_SECRET"),
	// }

	service := services.NewOneIDService()
	// data := service.LoginPWD(os.Getenv("USERNAME"), os.Getenv("PASSWORD"))

	t.Run("ShouldBeSuccess", func(t *testing.T) {

		response := service.GetCheckEmail("asdasdasd@gmail.com")

		if response.Result != "Success" {
			t.Error(response.Result)
			log.Println(response)
		}
	})

	t.Run("ShouldBeFail", func(t *testing.T) {
		response := service.GetCheckEmail("thinnawart@gmail.com")

		if response.Result == "Success" {
			t.Error(response.Result)
		}
	})
}
