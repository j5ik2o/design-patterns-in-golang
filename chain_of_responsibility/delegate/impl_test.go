package delegate

import (
	"desgin-patterns-in-golang/chain_of_responsibility"
	"desgin-patterns-in-golang/chain_of_responsibility/delegate/supports"
	"testing"
)

func TestChainOfResponsibility(t *testing.T) {
	fred := supports.NewLimitSupport("Fred", 300, nil)
	elmo := supports.NewOddSupport("Elmo", fred)
	diana := supports.NewLimitSupport("Diana", 200, elmo)
	charlie := supports.NewSpecialSupport("Charlie", 429, diana)
	bob := supports.NewLimitSupport("Bob", 100, charlie)
	alice := supports.NewNoSupport("Alice", bob)

	for i := 0; i < 500; i += 33 {
		alice.Support(chain_of_responsibility.NewTrouble(i))
	}
}
