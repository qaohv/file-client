package main

import "fmt"

type UnknowCommand struct {
	Text string
}

func NewUnknowCommand(input string) UnknowCommand {
	return UnknowCommand{
		Text: fmt.Sprintf("Command '%s' is not supported, try help for getting command list", input),
	}
}

func (unknowCommand UnknowCommand) Do() string {
	return unknowCommand.Text
}

func (unknowCommand UnknowCommand) isExitCommand() bool {
	return false
}
