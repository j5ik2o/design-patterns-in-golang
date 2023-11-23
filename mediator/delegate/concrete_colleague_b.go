package delegate

import "fmt"

type ConcreteColleagueB struct {
	mediator Mediator
	name     string
}

func NewConcreteColleagueB(mediator Mediator, name string) *ConcreteColleagueB {
	return &ConcreteColleagueB{
		mediator: mediator,
		name:     name,
	}
}

func (c *ConcreteColleagueB) GetName() string {
	return c.name
}

func (c *ConcreteColleagueB) OnChanged(msg string) {
	fmt.Printf("%s received: %s\n", c.name, msg)
}

func (c *ConcreteColleagueB) Run() {
	c.mediator.ColleagueChanged(c, "Hi!")
}
