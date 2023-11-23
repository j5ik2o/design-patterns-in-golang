package delegate

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
