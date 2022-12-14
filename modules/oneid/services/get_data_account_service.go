package services

import (
	"encoding/json"

	"github.com/benzdeus/oneplatform/env"
	"github.com/benzdeus/oneplatform/helper"
	one_entities "github.com/benzdeus/oneplatform/one_entities"
)

func (service *OneIDService) GetAccountData(accessToken string) one_entities.ResponseOneIDGetAccountData {

	response := one_entities.ResponseOneIDGetAccountData{}
	requestBody := map[string]interface{}{
		"client_id":     env.Options.ClientID,
		"client_secret": env.Options.ClientSecret,
	}

	body, statusCode := helper.RunOneIDHTTP("GET", env.BaseUrlOneID+"/account", requestBody, accessToken)

	response.ResponseCode = statusCode

	if body != nil {
		json.Unmarshal(body, &response)
	}

	return response
}
