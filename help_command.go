package main

type HelpCommand struct {
	Text string
}

const DEFAULT_HELP_TEXT = "Command list:\n help - show command list.\n list - show list files.\n get filename - download file and save this to specified folder.\n exit - exit programm."

func NewHelpCommand(helpText string) HelpCommand {
	return HelpCommand{
		Text: helpText,
	}
}

func (helpCommand HelpCommand) Do() string {
	return helpCommand.Text
}

func (helpCommand HelpCommand) isExitCommand() bool {
	return false
}
