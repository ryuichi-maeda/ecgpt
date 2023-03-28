package utils

import (
	"os"
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
