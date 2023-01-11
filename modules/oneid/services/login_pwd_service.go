package services

import (
	"encoding/json"

	"github.com/benzdeus/oneplatform/env"
	"github.com/benzdeus/oneplatform/helper"
	one_entities "github.com/benzdeus/oneplatform/one_entities"
)

func (service *OneIDService) LoginPWD(username, password string) one_entities.ResponseOneIDLoginPWD {

	response := one_entities.ResponseOneIDLoginPWD{}
	requestBody := map[string]interface{}{
		"username":      username,
		"password":      password,
		"grant_type":    "password",
		"client_id":     env.Options.ClientID,
		"client_secret": env.Options.ClientSecret,
	}
	body, statusCode := helper.RunOneIDHTTP("POST", env.BaseUrlOneID+"/oauth/getpwd", requestBody, "")

	if statusCode == 200 {
		response.ResponseCode = 200
	}

	if body != nil {
		json.Unmarshal(body, &response)
	}

	return response
}
