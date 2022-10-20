package mediator

import (
	"fmt"
	"github.com/igrmk/treemap/v2"
)

type Colleague interface {
	GetName() string
	OnChanged(msg string)
	Run()
}

type Mediator interface {
	AddColleague(colleague Colleague)
	ColleagueChanged(colleagueUpdated Colleague, msg string)
}

type ConcreteMediator struct {
	colleagues *treemap.TreeMap[string, Colleague]
}

func NewConcreteMediator() *ConcreteMediator {
	return &ConcreteMediator{
		colleagues: treemap.New[string, Colleague](),
	}
}

func (c *ConcreteMediator) AddColleague(colleague Colleague) {
	c.colleagues.Set(colleague.GetName(), colleague)
}

func (c *ConcreteMediator) ColleagueChanged(colleague Colleague, msg string) {
	for it := c.colleagues.Iterator(); it.Valid(); it.Next() {
		if it.Key() != colleague.GetName() {
			colleague.OnChanged(msg)
		}
	}
}

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
