package helper

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/benzdeus/oneplatform/env"
	one_entities "github.com/benzdeus/oneplatform/one_entities"
)

func RunOneChatSendComponentHTTP(method, titile, url, oneID, botID string, templates []one_entities.OneChatRequestElementSendComponent) ([]byte, int) {

	requestBody, _ := json.Marshal(map[string]interface{}{
		"to":                  []string{oneID},
		"bot_id":              botID,
		"type":                "template",
		"custom_notification": titile,
		"elements":            templates,
	})

	request, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))

	if method != "GET" && method != "DELETE" {
		request.Header.Set("Content-Type", "application/json")
	}

	if env.Options.OneChatAPIKey != "" {
		request.Header.Set("Authorization", "Bearer "+env.Options.OneChatAPIKey)
	}

	client := &http.Client{}

	response, errResponse := client.Do(request)
	if errResponse != nil {
		log.Println(errResponse)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	return body, response.StatusCode

}
