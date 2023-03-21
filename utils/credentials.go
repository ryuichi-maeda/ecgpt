package utils

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"ecgpt/structs"
)

func GetCredentials() (*structs.Credentials, error) {
	credentialsFilePath, err := GetCredentialsFilePath()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(credentialsFilePath)
	if err != nil {
		err = errors.New("error: Before running this command, OpenAI API key must be configured with 'ecgpt configure' command. ")
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var credentials structs.Credentials
	json.Unmarshal(data, &credentials)

	return &credentials, nil
}
