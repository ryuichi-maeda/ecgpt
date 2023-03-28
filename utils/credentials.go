package utils

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"ecgpt/config"
)

type Credentials struct {
	OpenAIAPIKey string `json:"open_ai_api_key"`
}

func GetCredentialsFilePath() (string, error) {
	configDir, err := GetConfigDirPath()
	if err != nil {
		return "", err
	}

	return configDir + "/" + config.CREDENTIALS_FILE, nil
}

func GetNewCredentialsFile() (*os.File, error) {
	credentialsFilePath, err := GetCredentialsFilePath()
	if err != nil {
		return nil, err
	}

	// If $HOME/.ecgpt/credentials.json exist, remove the file
	if _, err := os.Stat(credentialsFilePath); os.IsExist(err) {
		err := os.Remove(credentialsFilePath)
		if err != nil {
			return nil, err
		}
	}

	// Create a new $HOME/.ecgpt/credentials.json file
	file, err := os.Create(credentialsFilePath)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func GetCredentials() (*Credentials, error) {
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

	var credentials Credentials
	json.Unmarshal(data, &credentials)

	return &credentials, nil
}
