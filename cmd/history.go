/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"ecgpt/config"
	"ecgpt/utils"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

func getHistoryFilePaths() ([]string, error) {
	historyDirPath, err := utils.GetHistoryDirPath()
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

func getHistoryFilePath() (string, error) {
	historyFilePaths, err := getHistoryFilePaths()
	if err != nil {
		return "", err
	}

	prompt := promptui.Select{
		Label: "Select a history",
		Items: historyFilePaths,
	}

	_, historyFilePath, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return historyFilePath, err
}

// historyCmd represents the history command
var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Check past conversations",
	Long:  `You can check a past conversation you selected.`,
	Run: func(cmd *cobra.Command, args []string) {
		historyFilePath, err := getHistoryFilePath()
		if err != nil {
			fmt.Println(err)
			return
		}

		historyDirPath, err := utils.GetHistoryDirPath()
		if err != nil {
			fmt.Println(err)
			return
		}

		file, err := os.Open(historyDirPath + "/" + historyFilePath + ".json")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		data, err := io.ReadAll(file)
		if err != nil {
			fmt.Println(err)
			return
		}

		var msgs []openai.ChatCompletionMessage
		json.Unmarshal(data, &msgs)

		for _, msg := range msgs {
			switch msg.Role {
			case openai.ChatMessageRoleUser:
				fmt.Print(config.ROLE_OUTPUT_USER)
			case openai.ChatMessageRoleAssistant:
				fmt.Print(config.ROLE_OUTPUT_ASSISTANT)
			default:
				continue
			}

			fmt.Println(msg.Content)
		}
	},
}

func init() {
	rootCmd.AddCommand(historyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// historyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// historyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
