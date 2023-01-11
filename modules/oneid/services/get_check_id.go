package services

import (
	"encoding/json"

	"github.com/benzdeus/oneplatform/env"
	"github.com/benzdeus/oneplatform/helper"
	one_entities "github.com/benzdeus/oneplatform/one_entities"
)

func (service *OneIDService) GetCheckId(idCard string) one_entities.ResponseOneIdCheckId {

	response := one_entities.ResponseOneIdCheckId{}
	requestBody := map[string]interface{}{}

	body, statusCode := helper.RunOneIDHTTP("GET", env.BaseUrlOneID+"/check_id/?idCard="+idCard, requestBody, "")

	response.ResponseCode = statusCode

	if body != nil {
		json.Unmarshal(body, &response)
	}

	return response

}
