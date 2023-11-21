package embed

import (
	"desgin-patterns-in-golang/chain_of_responsibility"
)

type OddSupport struct {
	*SupportBase
}

func NewOddSupport(name string, next *SupportBase) *OddSupport {
	self := &OddSupport{}
	self.SupportBase = NewSupportBase(name, oddSupportResolve, next)
	return self
}

func oddSupportResolve(trouble *chain_of_responsibility.Trouble) bool {
	return trouble.GetNumber()%2 == 1
}
