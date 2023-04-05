package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestGetNewCredentialsFile(t *testing.T) {
	// Create a temporary home directory for testing
	tmpHomeDir, err := os.MkdirTemp("", "test_home")
	if err != nil {
		t.Fatalf("failed to create temporary home directory: %v", err)
	}
	defer os.RemoveAll(tmpHomeDir)

	// Set the temporary home directory as the user's home directory
	origHomeDir := os.Getenv("HOME")
	defer os.Setenv("HOME", origHomeDir)
	os.Setenv("HOME", tmpHomeDir)

	// Test the function when the credentials file does not exist
	file, err := GetNewCredentialsFile()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if _, err := os.Stat(tmpHomeDir + "/.ecgpt/credentials.json"); os.IsNotExist(err) {
		t.Errorf("expected credentials file to be created but not found: %v", err)
	}
	if file == nil {
		t.Errorf("expected file to be created but got nil")
	}
	file.Close()

	// Test the function when the credentials file already exists
	file, err = GetNewCredentialsFile()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if _, err := os.Stat(tmpHomeDir + "/.ecgpt/credentials.json"); os.IsNotExist(err) {
		t.Errorf("expected credentials file to exist but not found: %v", err)
	}
	if file == nil {
		t.Errorf("expected file to be created but got nil")
	}
	file.Close()
}

func TestGetCredentials(t *testing.T) {
	// Create a temporary home directory for testing
	tmpHomeDir, err := os.MkdirTemp("", "test_home")
	if err != nil {
		t.Fatalf("failed to create temporary home directory: %v", err)
	}
	defer os.RemoveAll(tmpHomeDir)

	// Set the temporary home directory as the user's home directory
	origHomeDir := os.Getenv("HOME")
	defer os.Setenv("HOME", origHomeDir)
	os.Setenv("HOME", tmpHomeDir)

	// Test the function when the credentials file does not exist
	_, err = GetCredentials()
	if err == nil {
		t.Errorf("expected error but got nil")
	}

	// Create a temporary credentials file for testing
	credentialsFilePath := tmpHomeDir + "/.ecgpt/credentials.json"
	err = os.MkdirAll(filepath.Dir(credentialsFilePath), 0755)
	if err != nil {
		t.Fatalf("failed to create temporary credentials file: %v", err)
	}
	credentialsFile, err := os.Create(credentialsFilePath)
	if err != nil {
		t.Fatalf("failed to create temporary credentials file: %v", err)
	}
	credentialsFile.Close()

	// Test the function when the credentials file exists
	savedCredentials := Credentials{
		OpenAIAPIKey: "test_openai_key",
	}
	credentialsBytes, err := json.Marshal(savedCredentials)
	if err != nil {
		t.Fatalf("failed to marshal credentials: %v", err)
	}
	err = os.WriteFile(credentialsFilePath, credentialsBytes, 0644)
	if err != nil {
		t.Fatalf("failed to write credentials to file: %v", err)
	}
	credentials, err := GetCredentials()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if credentials.OpenAIAPIKey != "test_openai_key" {
		t.Errorf("expected credentials.OpenAIKey to be 'test_openai_key' but got '%s'", credentials.OpenAIAPIKey)
	}
}
