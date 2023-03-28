package utils

import (
	"fmt"
	"os"
	"strings"
	"time"

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

func GetHistoryDirPath() (string, error) {
	configDir, err := GetConfigDirPath()
	if err != nil {
		return "", err
	}

	historyDirPath := configDir + "/" + config.HISTORY_DIR

	// If $HOME/.ecgpt dir does not exist, create the dir
	if f, err := os.Stat(historyDirPath); os.IsNotExist(err) || !f.IsDir() {
		err = os.Mkdir(historyDirPath, 0755)
		if err != nil {
			return "", err
		}
	}

	return historyDirPath, nil
}

func GetNewHistoryFile(summary string) (*os.File, error) {
	historyDir, err := GetHistoryDirPath()
	if err != nil {
		return nil, err
	}

	// Create a new $HOME/.ecgpt/history/2006-01-02_15:04:05_{summary}.json file
	historyFilePath := historyDir + "/" + time.Now().Format("2006-01-02_15:04:05") + "_" + strings.ReplaceAll(summary, " ", "_") + ".json"
	file, err := os.Create(historyFilePath)
	if err != nil {
		return nil, err
	}

	return file, nil
}
