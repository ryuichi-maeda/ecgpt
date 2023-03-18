package utils

import (
	"bytes"
	"fmt"
	"net/http"

	"ecgpt/config"
)

func CreateRequestForOpenAIAPI(jsonString []byte) (*http.Request, error) {
	req, err := http.NewRequest("POST", config.ENDPOINT, bytes.NewReader(jsonString))
	if err != nil {
		return nil, err
	}

	credentials, err := GetCredentials()
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", credentials.OpenAIAPIKey))

	return req, nil
}
