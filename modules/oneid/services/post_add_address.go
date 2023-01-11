package services

import (
	"encoding/json"

	"github.com/benzdeus/oneplatform/env"
	"github.com/benzdeus/oneplatform/helper"
	one_entities "github.com/benzdeus/oneplatform/one_entities"
)

func (service *OneIDService) PostAddAddress(house_no, room_no, floor, building_name, moo_ban, moo_no, soi, yaek, street, amphoe, tambon, zipcode, province, primary, access_token string) one_entities.ResponseOneIdAddAddress {

	response := one_entities.ResponseOneIdAddAddress{}
	requestBody := map[string]interface{}{
		"client_id":     env.Options.ClientID,
		"secret_key":    env.Options.ClientSecret,
		"house_no":      house_no,
		"room_no":       room_no,
		"floor":         floor,
		"buliding_name": building_name,
		"moo_ban":       moo_ban,
		"moo_no":        moo_no,
		"soi":           soi,
		"yaek":          yaek,
		"street":        street,
		"amphoe":        amphoe,
		"tambon":        tambon,
		"zipcode":       zipcode,
		"province":      province,
		"primary":       primary,
	}

	body, statusCode := helper.RunOneIDHTTP("POST", env.BaseUrlOneID+"/add_address", requestBody, access_token)

	response.ResponseCode = statusCode

	if body != nil {
		json.Unmarshal(body, &response)
	}

	return response

}
