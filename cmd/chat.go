/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"

	"ecgpt/config"
	"ecgpt/structs"
	"ecgpt/utils"
)

func createRequestBody(content string) structs.RequestBody {
	message := structs.Message{
			Role:    "user",
			Content: content,
	}
	messages := []structs.Message{message}
	reqBody := structs.RequestBody{
		Model:    config.AI_MODEL,
		Messages: messages,
	}

	return reqBody
}

func getUserMessage() string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("You       > ")
	scanner.Scan()

	return scanner.Text()
}

// chatCmd represents the chat command
var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Chat with AI assistant",
	Long: `You can chat with AI assistant via CLI interface.
Before running this command, OpenAI API key must be configured with 'ecgpt configure' command.`,
	Run: func(cmd *cobra.Command, args []string) {
		userMessage := getUserMessage()
		reqBody := createRequestBody(userMessage)

		jsonString, err := json.Marshal(reqBody)
		if err != nil {
			fmt.Println(err)
			return
		}

		req, err := utils.CreateRequestForOpenAIAPI(jsonString)
		if err != nil {
			fmt.Println(err)
			return
		}

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		// If the response is not OK
		if res.StatusCode != 200 {
			switch {
			case res.StatusCode == 401:
				fmt.Println("OpenAI API Key you set may be incorrect.")
			default:
				fmt.Println(res.Status)
			}
			return
		}

		// Decode the response body
		resBody := &structs.ResponseBody{}
		decoder := json.NewDecoder(res.Body)
		err = decoder.Decode(resBody)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, choice := range resBody.Choices {
			fmt.Println("Assistant > ", choice.Message.Content)
		}
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chatCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chatCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
