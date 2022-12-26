package services

import (
	"encoding/json"
	"log"

	"github.com/benzdeus/oneplatform/env"
	"github.com/benzdeus/oneplatform/helper"
	"github.com/benzdeus/oneplatform/one_entities"
)

func (servive OneIDService) Register(registerRequest one_entities.RequestOneIDRegister) one_entities.ResponseOneIDRegister {

	registerRequest.SetKey()
	log.Println(registerRequest.GetKey())
	body, statusCode := helper.RunOneIDHTTP("POST", env.BaseUrlOneID+"/citizen/register", registerRequest, "")

	response := one_entities.ResponseOneIDRegister{}
	if statusCode != 200 && statusCode != 400 {
		response.Code = statusCode
		return response
	}

	err := json.Unmarshal(body, &response)
	if err != nil {
		response.Code = 500
		response.ErrorMessage = map[string][]string{
			"system": {err.Error()},
		}
		return response
	}

	if response.Result == "Fail" {
		response.Code = 400
	}

	return response
}
