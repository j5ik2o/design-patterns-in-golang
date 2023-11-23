package delegate

import "github.com/igrmk/treemap/v2"

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
