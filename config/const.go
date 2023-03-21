package config

const (
	// Dir
	CONFIG_DIR  = ".ecgpt"
	HISTORY_DIR = "history"

	// File
	CREDENTIALS_FILE = "credentials.json"
	BEHAVIOR_FILE    = "behavior.txt"

	// Colored Role Output
	ROLE_OUTPUT_USER      = "\x1b[34m" + "You       > " + "\x1b[0m" // Blue
	ROLE_OUTPUT_ASSISTANT = "\x1b[35m" + "Assistant > " + "\x1b[0m" // Magenta
	ROLE_OUTPUT_SYSTEM    = "\x1b[31m" + "System    > " + "\x1b[0m" // Cyan
)
