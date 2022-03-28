package command

import "fmt"

type Command interface {
	Execute()
}

// --- MacroCommand

type MacroCommand struct {
	commands []Command
}

func NewMacroCommand() *MacroCommand {
	return &MacroCommand{
		commands: []Command{},
	}
}

func (mc *MacroCommand) Execute() {
	for _, c := range mc.commands {
		c.Execute()
	}
}

func (mc *MacroCommand) Append(cmd Command) *MacroCommand {
	return &MacroCommand{
		commands: append(mc.commands, cmd),
	}
}

func (mc *MacroCommand) Undo() *MacroCommand {
	return &MacroCommand{
		commands: mc.commands[:len(mc.commands)-1],
	}
}

func (mc *MacroCommand) Clear() *MacroCommand {
	return NewMacroCommand()
}

// --- EchoCommand

type EchoCommand struct {
	msg string
}

func NewEchoCommand(msg string) *EchoCommand {
	return &EchoCommand{
		msg: msg,
	}
}

func (mc *EchoCommand) Execute() {
	fmt.Println(mc.msg)
}
