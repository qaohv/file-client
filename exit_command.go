package main

type ExitCommand struct {
	ExitMessage string
}

func NewExitCommand() ExitCommand {
	return ExitCommand{
		ExitMessage: "Session closed, bye!",
	}
}

func (exitCommand ExitCommand) Do() string {
	return exitCommand.ExitMessage
}

func (exitCommand ExitCommand) isExitCommand() bool {
	return true
}
