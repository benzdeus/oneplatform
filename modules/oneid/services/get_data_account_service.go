package services

import (
	"encoding/json"
	"github.com/benzdeus/oneplatform/entities"
	"github.com/benzdeus/oneplatform/env"
	"github.com/benzdeus/oneplatform/helper"
)

func (service *OneIDService) GetAccountData(accessToken string) entities.ResponseOneIDGetAccountData {

	response := entities.ResponseOneIDGetAccountData{}
	requestBody := map[string]interface{}{
		"client_id":     env.Options.ClientID,
		"client_secret": env.Options.ClientSecret,
	}

	body, statusCode := helper.SetRequestOneID("GET", env.BaseUrlOneID+"/account", requestBody, accessToken)

	if statusCode == 200 {
		response.ResponseCode = 200
	}

	if body != nil {
		json.Unmarshal(body, &response)
	}

	return response
}
