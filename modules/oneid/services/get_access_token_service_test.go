package services_test

import (
	"testing"

	"github.com/benzdeus/oneplatform/configs"
	"github.com/benzdeus/oneplatform/env"
	"github.com/benzdeus/oneplatform/modules/oneid/services"
)

func TestOneIDService_GetAccessToken(t *testing.T) {

	env.Options = configs.Options{
		ClientIDSharedToken:      "",
		ClientSecretSharedToken:  "",
		ClientRefCodeSharedToken: "",
	}

	sharedToken := ``

	oneidService := services.NewOneIDService()
	res := oneidService.GetAccessToken(sharedToken)

	// t.Errorf("%#+v", env.Options)
	t.Errorf("%#v", res)

}
