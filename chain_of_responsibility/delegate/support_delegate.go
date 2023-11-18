package delegate

import (
	"desgin-patterns-in-golang/chain_of_responsibility"
	"fmt"
)

type SupportDelegate struct {
	name     string
	resolver Resolver
	next     *SupportDelegate
}

func NewSupportDelegate(name string, resolver Resolver, next *SupportDelegate) *SupportDelegate {
	return &SupportDelegate{
		name,
		resolver,
		next,
	}
}

func (s *SupportDelegate) Support(trouble *chain_of_responsibility.Trouble) {
	if s.resolver(trouble) {
		s.done(trouble)
	} else if s.next != nil {
		s.next.Support(trouble)
	} else {
		s.fail(trouble)
	}
}

func (s *SupportDelegate) done(trouble *chain_of_responsibility.Trouble) {
	fmt.Printf("%s is resolved by %s.\n", trouble, s.name)
}

func (s *SupportDelegate) fail(trouble *chain_of_responsibility.Trouble) {
	fmt.Printf("%s cannot be resolved.\n", trouble)
}
