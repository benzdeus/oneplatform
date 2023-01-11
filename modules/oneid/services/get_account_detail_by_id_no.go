package services

import (
	"encoding/json"

	"github.com/benzdeus/oneplatform/env"
	"github.com/benzdeus/oneplatform/helper"
	one_entities "github.com/benzdeus/oneplatform/one_entities"
)

func (service *OneIDService) GetAccountDetialByIdNo(id_no, access_token string) one_entities.ResponseOneIdAccountDetailByIdNo {

	response := one_entities.ResponseOneIdAccountDetailByIdNo{}
	requestBody := map[string]interface{}{
		"client_id": env.Options.ClientID,
		"secretKey": env.Options.ClientSecret,
	}

	body, statusCode := helper.RunOneIDHTTP("GET", env.BaseUrlOneID+"/get_account_detail_by_id_no/"+id_no, requestBody, access_token)

	response.Code = statusCode

	if body != nil {
		json.Unmarshal(body, &response)
	}

	return response

}
