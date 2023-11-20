package supports

import (
	"desgin-patterns-in-golang/chain_of_responsibility"
	delegate2 "desgin-patterns-in-golang/chain_of_responsibility/delegate"
	"fmt"
)

type SpecialSupport struct {
	name     string
	number   int
	delegate *delegate2.SupportDelegate
}

func NewSpecialSupport(name string, number int, next delegate2.Support) *SpecialSupport {
	var nextDelegate *delegate2.SupportDelegate = nil
	if next != nil {
		nextDelegate = next.GetDelegate()
	}
	self := &SpecialSupport{
		name:   name,
		number: number,
	}
	self.delegate = delegate2.NewSupportDelegate(name, self.specialResolver, nextDelegate)
	return self
}

func (s *SpecialSupport) specialResolver(trouble *chain_of_responsibility.Trouble) bool {
	if trouble.GetNumber() == s.number {
		return true
	}
	return false
}

func (s *SpecialSupport) Support(trouble *chain_of_responsibility.Trouble) {
	s.delegate.Support(trouble)
}

func (s *SpecialSupport) GetDelegate() *delegate2.SupportDelegate {
	return s.delegate
}

func (s *SpecialSupport) String() string {
	return fmt.Sprintf("[%s@SpecialSupport]", s.name)
}
