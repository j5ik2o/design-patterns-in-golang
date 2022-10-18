package state

import "testing"

func TestState(t *testing.T) {
	state := Day{}
	context := NewStateContext(&state)
	context.Run()
}
