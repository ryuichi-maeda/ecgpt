package utils

import (
	"fmt"
	"os"

	"ecgpt/config"
)

func GetConfigDirPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	configDirPath := homeDir + "/" + config.CONFIG_DIR

	// If $HOME/.ecgpt dir does not exist, create the dir
	if f, err := os.Stat(configDirPath); os.IsNotExist(err) || !f.IsDir() {
		err = os.Mkdir(configDirPath, 0755)
		if err != nil {
			return "", err
		}
	}

	return configDirPath, nil
}

func GetCredentialsFile() (*os.File, error) {
	configDir, err := GetConfigDirPath()
	if err != nil {
		return nil, err
	}

	credentialsFilePath := configDir + "/" + config.CREDENTIALS_FILE

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
