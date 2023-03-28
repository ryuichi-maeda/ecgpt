package utils

import (
	"os"
	"testing"
)

func TestGetConfigDirPath(t *testing.T) {
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

	// Test the function
	path, err := GetConfigDirPath()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if path != tmpHomeDir+"/.ecgpt" {
		t.Errorf("unexpected path: %s", path)
	}
}
