package services

import (
	"encoding/json"

	"github.com/benzdeus/oneplatform/env"
	"github.com/benzdeus/oneplatform/helper"
	one_entities "github.com/benzdeus/oneplatform/one_entities"
	"github.com/dgrijalva/jwt-go"
)

type sharedToken struct {
	jwt.StandardClaims
	AccessToken string `json:"access_token"`
}

func (service OneIDService) GetAccessToken(st string) one_entities.ReponseONEIDGetAccessToken {
	dataSharedToken := new(sharedToken)
	response := one_entities.ReponseONEIDGetAccessToken{}

	_, errParseToken := jwt.ParseWithClaims(st, dataSharedToken, nil)

	if errParseToken == nil {
		response.Result = "Fail | Decode Stage | " + errParseToken.Error()
		return response
	}

	requestBody := map[string]interface{}{
		"client_id":     env.Options.ClientIDSharedToken,
		"client_secret": env.Options.ClientSecretSharedToken,
		"shared_token":  dataSharedToken.AccessToken,
		"refcode":       env.Options.ClientRefCodeSharedToken,
	}

	body, statusCode := helper.RunOneIDHTTP("POST", env.BaseUrlOneID+"/oauth/shared-token", requestBody, "")

	if statusCode != 200 {
		response.Result = "Fail | " + string(body)
		return response
	}

	if body != nil {
		err := json.Unmarshal(body, &response)
		if err != nil {
			response.Result = "Fail | Unmarshal Stage"
			return response
		}
	}

	return response
}
