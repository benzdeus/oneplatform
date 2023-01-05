package helper

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func RunOneIDHTTP(method string, url string, requestBodyJSON interface{}, acessToken string) ([]byte, int) {
	var requestBody []byte
	requestBody, _ = json.Marshal(requestBodyJSON)

	request, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))

	if method != "GET" && method != "DELETE" {
		request.Header.Set("Content-Type", "application/json")
	}

	if acessToken != "" {
		request.Header.Set("Authorization", "Bearer "+acessToken)
		request.Header.Set("client_id", os.Getenv("CLIENT_ID"))
		request.Header.Set("secretKey", os.Getenv("CLIENT_SECRET"))
		request.Header.Set("ref_code", "")
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
