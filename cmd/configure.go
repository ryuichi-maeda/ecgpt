/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	"ecgpt/utils"
)

func getEnteredApiKey() (string, error) {
	prompt := promptui.Prompt{
		Label: "OpenAI API Key",
	}
	apiKey, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return apiKey, nil
}

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Set a OpenAI API Key",
	Long:  `Set a OpenAI API key. You can get API keys from https://platform.openai.com/account/api-keys, but an OpenAI account is required.`,
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, err := getEnteredApiKey()
		if err != nil {
			fmt.Println(err)
			return
		}

		file, err := utils.GetNewCredentialsFile()
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		// Set credentials
		credentials := utils.Credentials{
			OpenAIAPIKey: apiKey,
		}

		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")
		encoder.Encode(credentials)
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configureCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configureCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
