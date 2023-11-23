package delegate

import "fmt"

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
