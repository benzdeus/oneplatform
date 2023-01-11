package services

import (
	"encoding/json"

	"github.com/benzdeus/oneplatform/env"
	"github.com/benzdeus/oneplatform/helper"
	"github.com/benzdeus/oneplatform/one_entities"
)

func (service *OneIDService) PostCheckOneMail(onemail string) one_entities.ResponseOneIDCheckOneMail {

	response := one_entities.ResponseOneIDCheckOneMail{}
	requestBody := map[string]interface{}{
		"client_id":  env.Options.ClientID,
		"secret_key": env.Options.ClientSecret,
		"onemail":    onemail,
	}

	body, statusCode := helper.RunOneIDHTTP("POST", env.BaseUrlOneID+"/check_one_mail", requestBody, "")

	response.ResponseCode = statusCode

	if body != nil {
		json.Unmarshal(body, &response)
	}

	return response

}
