package embed

import (
	"desgin-patterns-in-golang/chain_of_responsibility"
	"testing"
)

func TestChainOfResponsibility(t *testing.T) {
	fred := NewLimitSupport("Fred", 300, nil)
	elmo := NewOddSupport("Elmo", fred.SupportBase)
	diana := NewLimitSupport("Diana", 200, elmo.SupportBase)
	charlie := NewSpecialSupport("Charlie", 429, diana.SupportBase)
	bob := NewLimitSupport("Bob", 100, charlie.SupportBase)
	alice := NewNoSupport("Alice", bob.SupportBase)

	for i := 0; i < 500; i += 33 {
		alice.Support(chain_of_responsibility.NewTrouble(i))
	}
}
