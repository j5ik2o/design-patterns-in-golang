package chain_of_responsibility

import (
	"fmt"
	"github.com/samber/mo"
)

type SpecialSupport struct {
	name   string
	next   mo.Option[Support]
	number int
}

func NewSpecialSupport(name string, number int, next mo.Option[Support]) *SpecialSupport {
	return &SpecialSupport{
		name,
		next,
		number,
	}
}

func (s *SpecialSupport) String() string {
	return fmt.Sprintf("[%s@SpecialSupport]", s.name)
}

func (s *SpecialSupport) Done(trouble *Trouble) {
	fmt.Printf("%s is resolved by %s.\n", trouble, s.String())
}

func (s *SpecialSupport) Fail(trouble *Trouble) {
	fmt.Printf("%s cannot be resolved.\n", trouble)
}

func (s *SpecialSupport) Next() mo.Option[Support] {
	return s.next
}

func (s *SpecialSupport) Resolve(trouble *Trouble) bool {
	if trouble.GetNumber() == s.number {
		return true
	}
	return false
}

func (s *SpecialSupport) Support(trouble *Trouble) {
	support(s, trouble)
}
