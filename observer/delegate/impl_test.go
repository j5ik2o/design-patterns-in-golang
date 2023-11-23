package delegate

import (
	"testing"
)

func TestObserver(t *testing.T) {
	generator := NewRandomNumberGenerator()
	observer1 := DigitObserver{}
	observer2 := GraphObserver{}
	generator.AddObserver(&observer1)
	generator.AddObserver(&observer2)
	generator.Execute()
}
