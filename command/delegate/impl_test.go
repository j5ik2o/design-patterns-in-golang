package delegate

import "testing"

func executeCommand(cmd Command) {
	cmd.Execute()
}

func TestCommand(t *testing.T) {
	mc := NewMacroCommand()
	mc = mc.Append(NewEchoCommand("1"))
	mc = mc.Append(NewEchoCommand("2"))
	mc = mc.Append(NewEchoCommand("3"))

	echo := NewEchoCommand("0")
	cmds := []Command{mc, echo}

	for _, c := range cmds {
		executeCommand(c)
	}
}
