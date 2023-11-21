package embed

import (
	"desgin-patterns-in-golang/chain_of_responsibility"
)

type SpecialSupport struct {
	*SupportBase
	number int
}

func NewSpecialSupport(name string, number int, next *SupportBase) *SpecialSupport {
	self := &SpecialSupport{
		number: number,
	}
	self.SupportBase = NewSupportBase(name, self.resolve, next)
	return self
}

func (s *SpecialSupport) resolve(trouble *chain_of_responsibility.Trouble) bool {
	return trouble.GetNumber() == s.number
}
