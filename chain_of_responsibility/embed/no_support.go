package embed

import (
	"desgin-patterns-in-golang/chain_of_responsibility"
)

type NoSupport struct {
	*SupportBase
}

func NewNoSupport(name string, next *SupportBase) *NoSupport {
	self := &NoSupport{}
	self.SupportBase = NewSupportBase(name, noSupportResolve, next)
	return self
}

func noSupportResolve(trouble *chain_of_responsibility.Trouble) bool {
	return false
}
