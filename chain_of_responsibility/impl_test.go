package chain_of_responsibility

import (
	"github.com/samber/mo"
	"testing"
)

func TestChainOfResponsibility(t *testing.T) {
	fred := NewLimitSupport("Fred", 300, mo.None[Support]())
	elmo := NewOddSupport("Elmo", mo.Some[Support](fred))
	diana := NewLimitSupport("Diana", 200, mo.Some[Support](elmo))
	charlie := NewSpecialSupport("Charlie", 429, mo.Some[Support](diana))
	bob := NewLimitSupport("Bob", 100, mo.Some[Support](charlie))
	alice := NewNoSupport("Alice", mo.Some[Support](bob))

	for i := 0; i < 500; i += 33 {
		alice.Support(NewTrouble(i))
	}
}
