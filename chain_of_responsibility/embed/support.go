package embed

import (
	"desgin-patterns-in-golang/chain_of_responsibility"
	"fmt"
)

type ResolveF func(trouble *chain_of_responsibility.Trouble) bool

type SupportBase struct {
	name     string
	resolver ResolveF
	next     *SupportBase
}

func NewSupportBase(name string, resolver ResolveF, next *SupportBase) *SupportBase {
	return &SupportBase{
		name,
		resolver,
		next,
	}
}

func (s *SupportBase) Support(trouble *chain_of_responsibility.Trouble) {
	if s.resolver(trouble) {
		s.done(trouble)
	} else if s.next != nil {
		s.next.Support(trouble)
	} else {
		s.fail(trouble)
	}
}

func (s *SupportBase) done(trouble *chain_of_responsibility.Trouble) {
	fmt.Printf("%s is resolved by %s.\n", trouble, s.name)
}

func (s *SupportBase) fail(trouble *chain_of_responsibility.Trouble) {
	fmt.Printf("%s cannot be resolved.\n", trouble)
}
