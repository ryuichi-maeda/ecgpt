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
