package services_test

import (
	"testing"

	"github.com/benzdeus/oneplatform/configs"
	"github.com/benzdeus/oneplatform/env"
	"github.com/benzdeus/oneplatform/modules/oneid/services"
	"github.com/benzdeus/oneplatform/one_entities"
)

func TestRegisterService(t *testing.T) {

	env.Options = configs.Options{
		ClientID:     "",
		ClientSecret: "",
		RefCode:      "",
	}

	oneIDService := services.NewOneIDService()
	registerRequest := one_entities.RequestOneIDRegister{
		AccountTitleTh:  "นาย",
		AccountTitleEng: "Mr",
		FirstNameTh:     "สามาน",
		FirstNameEng:    "saman",
		LastNameTh:      "กกกกกกกก",
		LastNameEng:     "kkkkkkkkkk",
		IdCardType:      "ID_CARD",
		IdCardNum:       "4881917526348",
		Email:           "sakjdksadjsakd@gmail.com",
		MobileNo:        "0999999999",
		BirthDate:       "1995-10-02",
		Username:        "bbkkjsjsjss",
		Password:        "V22344489s",
	}
	res := oneIDService.Register(registerRequest)
	t.Errorf("%#+v", res)

}
