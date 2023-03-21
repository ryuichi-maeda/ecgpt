/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	openai "github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"

	"ecgpt/config"
	"ecgpt/utils"
)

func getUserMsg() string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(config.ROLE_OUTPUT_USER)
	scanner.Scan()

	return scanner.Text()
}

func chatCompletion(client openai.Client, ctx context.Context, req openai.ChatCompletionRequest) (string, error) {
	var resMsg string

	stream, err := client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		return "", err
	}
	defer stream.Close()

	fmt.Print(config.ROLE_OUTPUT_ASSISTANT)
	for {
		receivedResponse, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("")
			return resMsg, nil
		}

		if err != nil {
			fmt.Println(err)
			return "", err
		}
		fmt.Printf("%s", receivedResponse.Choices[0].Delta.Content)
		resMsg += receivedResponse.Choices[0].Delta.Content
	}
}

func getBehaviorContent() (string, error) {
	content, err := os.ReadFile(config.BEHAVIOR_FILE)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func addReqMsg(role string, content string, reqMsgs *[]openai.ChatCompletionMessage) []openai.ChatCompletionMessage {
	msg := openai.ChatCompletionMessage{
		Role:    role,
		Content: content,
	}
	return append(*reqMsgs, msg)
}

func saveConversation(chatCompletionMessage *[]openai.ChatCompletionMessage) error {
	file, err := utils.GetNewHistoryFile()
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	encoder.Encode(chatCompletionMessage)
	return nil
}

// chatCmd represents the chat command
var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Chat with AI assistant",
	Long: `You can chat with AI assistant via CLI interface.
Before running this command, OpenAI API key must be configured with 'ecgpt configure' command.`,
	Run: func(cmd *cobra.Command, args []string) {
		var reqMsgs []openai.ChatCompletionMessage

		credentials, err := utils.GetCredentials()
		if err != nil {
			fmt.Println(err)
			return
		}
		client := openai.NewClient(credentials.OpenAIAPIKey)
		ctx := context.Background()

		isFirst := true
		for {
			var (
				role    string
				content string
			)

			// If the chat is first turn, set the behavior of the assistant
			if isFirst {
				role = openai.ChatMessageRoleSystem
				content, err = getBehaviorContent()
				if err != nil {
					fmt.Println(err)
					return
				}
				isFirst = false
			} else {
				role = openai.ChatMessageRoleUser
				content = getUserMsg()

				// Exit
				if content == "exit" {
					// Save conversation
					err := saveConversation(&reqMsgs)
					if err != nil {
						fmt.Println(err)
						return
					}
					break
				}
			}

			reqMsgs = addReqMsg(role, content, &reqMsgs)

			request := openai.ChatCompletionRequest{
				Model:    openai.GPT3Dot5Turbo,
				Messages: reqMsgs,
				Stream:   true,
			}

			resMsg, err := chatCompletion(*client, ctx, request)
			if err != nil {
				fmt.Println(err)
				return
			}

			reqMsgs = addReqMsg(openai.ChatMessageRoleAssistant, resMsg, &reqMsgs)
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
