package services

import (
	"encoding/json"

	"github.com/benzdeus/oneplatform/entities"
	"github.com/benzdeus/oneplatform/env"
	"github.com/benzdeus/oneplatform/helper"
)

func (service *OneIDService) LoginPWD(username, password string) entities.ResponseOneIDLoginPWD {

	response := entities.ResponseOneIDLoginPWD{}
	rquestBody := map[string]interface{}{
		"username":      username,
		"password":      password,
		"grant_type":    "password",
		"client_id":     env.Options.ClientID,
		"client_secret": env.Options.ClientSecret,
	}
	body, statusCode := helper.SetRequestOneID("POST", env.BaseUrl+"/oauth/getpwd", rquestBody, "")

	if statusCode == 200 {
		response.ResponseCode = 200
	}

	if body != nil {
		json.Unmarshal(body, &response)
	}

	return response
}
