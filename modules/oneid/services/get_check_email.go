package services

import (
	"encoding/json"

	"github.com/benzdeus/oneplatform/env"
	"github.com/benzdeus/oneplatform/helper"
	one_entities "github.com/benzdeus/oneplatform/one_entities"
)

func (service *OneIDService) GetCheckEmail(email string) one_entities.ResponseOneIDCheckEmail {

	response := one_entities.ResponseOneIDCheckEmail{}
	requestBody := map[string]interface{}{
		// "client_id":  env.Options.ClientID,
		// "secret_key": env.Options.ClientSecret,
		"email": email,
	}

	body, statusCode := helper.RunOneIDHTTP("GET", env.BaseUrlOneID+"/check_email/?email="+email, requestBody, "")

	response.ResponseCode = statusCode

	if body != nil {
		json.Unmarshal(body, &response)
	}

	return response

}
