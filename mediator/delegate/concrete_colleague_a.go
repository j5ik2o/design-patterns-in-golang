package delegate

import "fmt"

type ConcreteColleagueA struct {
	mediator Mediator
	name     string
}

func NewConcreteColleagueA(mediator Mediator, name string) *ConcreteColleagueA {
	return &ConcreteColleagueA{
		mediator: mediator,
		name:     name,
	}
}

func (c *ConcreteColleagueA) GetName() string {
	return c.name
}

func (c *ConcreteColleagueA) OnChanged(msg string) {
	fmt.Printf("%s received: %s\n", c.name, msg)
}

func (c *ConcreteColleagueA) Run() {
	c.mediator.ColleagueChanged(c, "Hello!")
}
