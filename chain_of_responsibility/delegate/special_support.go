package delegate

import (
	"desgin-patterns-in-golang/chain_of_responsibility"
	"fmt"
)

type SpecialSupport struct {
	name     string
	number   int
	delegate *SupportDelegate
}

func NewSpecialSupport(name string, number int, next Support) *SpecialSupport {
	var nextDelegate *SupportDelegate = nil
	if next != nil {
		nextDelegate = next.GetDelegate()
	}
	self := &SpecialSupport{
		name:   name,
		number: number,
	}
	self.delegate = NewSupportDelegate(name, self.specialResolver, nextDelegate)
	return self
}

func (s *SpecialSupport) specialResolver(trouble *chain_of_responsibility.Trouble) bool {
	return trouble.GetNumber() == s.number
}

func (s *SpecialSupport) Support(trouble *chain_of_responsibility.Trouble) {
	s.delegate.Support(trouble)
}

func (s *SpecialSupport) GetDelegate() *SupportDelegate {
	return s.delegate
}

func (s *SpecialSupport) String() string {
	return fmt.Sprintf("[%s@SpecialSupport]", s.name)
}
