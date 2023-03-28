package config

const (
	// Dir
	CONFIG_DIR  = ".ecgpt"
	HISTORY_DIR = "history"

	// File
	CREDENTIALS_FILE = "credentials.json"

	// Behavior
	BEHAVIOR_CONTENT = `
# Instructions
You are an American professional English conversation teacher.
Please chat with me under the following constrains.

# Constrains
If there is any room for improvement in my English sentence, you advise how to improve it.
You give a talk thema at the begginning and wait for response.
Also, you choose the talk thema at random so that the same talk thema does not continue.`

	// Colored Role Output
	ROLE_OUTPUT_USER      = "\x1b[34m" + "You       > " + "\x1b[0m" // Blue
	ROLE_OUTPUT_ASSISTANT = "\x1b[35m" + "Assistant > " + "\x1b[0m" // Magenta
	ROLE_OUTPUT_SYSTEM    = "\x1b[31m" + "System    > " + "\x1b[0m" // Cyan
)
