package services

import (
	"encoding/json"

	one_entities "github.com/benzdeus/oneplatform/entities"
	"github.com/benzdeus/oneplatform/env"
	"github.com/benzdeus/oneplatform/helper"
)

func (service *OneIDService) LoginPWD(username, password string) one_entities.ResponseOneIDLoginPWD {

	response := one_entities.ResponseOneIDLoginPWD{}
	rquestBody := map[string]interface{}{
		"username":      username,
		"password":      password,
		"grant_type":    "password",
		"client_id":     env.Options.ClientID,
		"client_secret": env.Options.ClientSecret,
	}
	body, statusCode := helper.RunOneIDHTTP("POST", env.BaseUrlOneID+"/oauth/getpwd", rquestBody, "")

	if statusCode == 200 {
		response.ResponseCode = 200
	}

	if body != nil {
		json.Unmarshal(body, &response)
	}

	return response
}
