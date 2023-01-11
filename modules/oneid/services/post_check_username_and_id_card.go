package services

import (
	"encoding/json"

	"github.com/benzdeus/oneplatform/env"
	"github.com/benzdeus/oneplatform/helper"
	one_entities "github.com/benzdeus/oneplatform/one_entities"
)

func (service *OneIDService) PostCheckUsernameAndIdCard(username, id_card string) one_entities.ResponseOneIdCheckUsernameAndIdCard {

	response := one_entities.ResponseOneIdCheckUsernameAndIdCard{}
	requestBody := map[string]interface{}{
		"client_id":  env.Options.ClientID,
		"secret_key": env.Options.ClientSecret,
		"username":   username,
		"id_card":    id_card,
	}

	body, statusCode := helper.RunOneIDHTTP("POST", env.BaseUrlOneID+"/check_username_and_id_card", requestBody, "")

	response.ResponseCode = statusCode

	if body != nil {
		json.Unmarshal(body, &response)
	}

	return response

}
