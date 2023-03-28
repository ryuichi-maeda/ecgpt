package utils

import (
	"os"
	"testing"
	"time"
)

func TestGetHistoryDirPath(t *testing.T) {
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

	// Test the function when the history directory does not exist
	dirPath, err := GetHistoryDirPath()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if _, err := os.Stat(tmpHomeDir + "/.ecgpt/history"); os.IsNotExist(err) {
		t.Errorf("expected history directory to be created but not found: %v", err)
	}
	if dirPath == "" {
		t.Errorf("expected directory path to be created but got empty string")
	}

	// Test the function when the history directory already exists
	dirPath, err = GetHistoryDirPath()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if _, err := os.Stat(tmpHomeDir + "/.ecgpt/history"); os.IsNotExist(err) {
		t.Errorf("expected history directory to exist but not found: %v", err)
	}
	if dirPath == "" {
		t.Errorf("expected directory path to be created but got empty string")
	}
}

func TestGetNewHistoryFile(t *testing.T) {
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

	summary := "summary"

	// Test the function when the history file does not exist
	file, err := GetNewHistoryFile(summary)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if _, err := os.Stat(tmpHomeDir + "/.ecgpt/history/" + time.Now().Format("2006-01-02_15:04:05") + "_" + summary + ".json"); os.IsNotExist(err) {
		t.Errorf("expected history file to be created but not found: %v", err)
	}
	if file == nil {
		t.Errorf("expected file to be created but got nil")
	}
	file.Close()

	// Test the function when the history file already exists
	file, err = GetNewHistoryFile(summary)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if _, err := os.Stat(tmpHomeDir + "/.ecgpt/history/" + time.Now().Format("2006-01-02_15:04:05") + "_" + summary + ".json"); os.IsNotExist(err) {
		t.Errorf("expected history file to exist but not found: %v", err)
	}
	if file == nil {
		t.Errorf("expected file to be created but got nil")
	}
	file.Close()
}
