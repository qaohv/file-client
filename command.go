package main

import (
	"strings"
)

const (
	LIST_FILES_COMMAND = "list"
	GET_FILE_COMMAND   = "get"
	HELP_COMMAND       = "help"
	EXIT_COMMAND       = "exit"
)

type Command interface {
	Do() string
	isExitCommand() bool
}

func makeCommand(input string, config *Config) Command {
	normalizedInput := strings.TrimSpace(input)

	inputParts := strings.Split(normalizedInput, " ")

	switch inputParts[0] {
	case HELP_COMMAND:
		return NewHelpCommand(DEFAULT_HELP_TEXT)
	case LIST_FILES_COMMAND:
		return NewListFilesCommand(config)
	case GET_FILE_COMMAND:
		if len(inputParts) > 1 {
			return NewGetFileCommand(strings.TrimSpace(inputParts[1]), config)
		} else {
			return NewHelpCommand(NO_FILE_ARGUMENT_MESSAGE)
		}
	case EXIT_COMMAND:
		return NewExitCommand()
	default:
		return NewUnknowCommand(inputParts[0])
	}
}
