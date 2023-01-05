package services

import (
	"encoding/json"

	"github.com/benzdeus/oneplatform/env"
	"github.com/benzdeus/oneplatform/helper"
	one_entities "github.com/benzdeus/oneplatform/one_entities"
)

func (service *OneIDService) PostDeleteEmail(email, access_token string) one_entities.ResponseOneIDDeleteEmail {
	response := one_entities.ResponseOneIDDeleteEmail{}
	requestBody := map[string]interface{}{
		"email":     email,
		"client_id": env.Options.ClientID,
		"secretKey": env.Options.ClientSecret,
	}

	body, statusCode := helper.RunOneIDHTTP("POST", env.BaseUrlOneID+"/delete_email", requestBody, access_token)

	response.ResponseCode = statusCode

	if body != nil {
		json.Unmarshal(body, &response)
	}

	return response

}
