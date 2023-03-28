package utils

import (
	"ecgpt/config"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/sashabaranov/go-openai"
)

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

func GetHistoryFilePaths() ([]string, error) {
	historyDirPath, err := GetHistoryDirPath()
	if err != nil {
		return nil, err
	}

	var historyFilePaths []string
	err = filepath.Walk(historyDirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		for i := len(path) - 1; i >= 0; i-- {
			if path[i] == '.' && path[i:] == ".json" {
				historyFilePaths = append(historyFilePaths, path[len(historyDirPath)+1:i])
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return historyFilePaths, nil
}

func SaveConversation(summary string, chatCompletionMessage *[]openai.ChatCompletionMessage) error {
	file, err := GetNewHistoryFile(summary)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	encoder.Encode(chatCompletionMessage)
	return nil
}
