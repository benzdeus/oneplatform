package services

import (
	"encoding/json"

	"github.com/benzdeus/oneplatform/env"
	"github.com/benzdeus/oneplatform/helper"
	one_entities "github.com/benzdeus/oneplatform/one_entities"
)

func (service *OneIDService) PostCheckUsernameAndMobile(username, mobile_no string) one_entities.ResponseOneIDCheckUsernameAndMobile {

	response := one_entities.ResponseOneIDCheckUsernameAndMobile{}
	requestBody := map[string]interface{}{
		"client_id":  env.Options.ClientID,
		"secret_key": env.Options.ClientSecret,
		"username":   username,
		"mobile_no":  mobile_no,
	}

	body, statusCode := helper.RunOneIDHTTP("POST", env.BaseUrlOneID+"/check_username_and_mobile", requestBody, "")

	response.ResponseCode = statusCode

	if body != nil {
		json.Unmarshal(body, &response)
	}

	return response

}
