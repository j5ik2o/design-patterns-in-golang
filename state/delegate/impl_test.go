package delegate

import "testing"

func TestState(t *testing.T) {
	context := NewStateContext(NewDay())
	context.Run()
}
