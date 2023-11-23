package delegate

import (
	"testing"
)

func TestMediator(t *testing.T) {

	mediator := NewConcreteMediator()
	colleagueA := NewConcreteColleagueA(mediator, "A")
	colleagueB := NewConcreteColleagueB(mediator, "B")

	mediator.AddColleague(colleagueA)
	mediator.AddColleague(colleagueB)

	colleagueA.Run()
	colleagueB.Run()
}
