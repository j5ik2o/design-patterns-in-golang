package delegate

import (
	"desgin-patterns-in-golang/chain_of_responsibility"
	"testing"
)

func TestChainOfResponsibility(t *testing.T) {
	fred := NewLimitSupport("Fred", 300, nil)
	elmo := NewOddSupport("Elmo", fred)
	diana := NewLimitSupport("Diana", 200, elmo)
	charlie := NewSpecialSupport("Charlie", 429, diana)
	bob := NewLimitSupport("Bob", 100, charlie)
	alice := NewNoSupport("Alice", bob)

	for i := 0; i < 500; i += 33 {
		alice.Support(chain_of_responsibility.NewTrouble(i))
	}
}
