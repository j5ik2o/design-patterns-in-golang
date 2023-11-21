package delegate

import (
	"desgin-patterns-in-golang/chain_of_responsibility"
	"fmt"
)

type NoSupport struct {
	name     string
	delegate *SupportDelegate
}

func NewNoSupport(name string, next Support) *NoSupport {
	var nextDelegate *SupportDelegate = nil
	if next != nil {
		nextDelegate = next.GetDelegate()
	}
	return &NoSupport{
		name,
		NewSupportDelegate(name, resolver, nextDelegate),
	}
}

func resolver(t *chain_of_responsibility.Trouble) bool {
	return false
}

func (n *NoSupport) Support(trouble *chain_of_responsibility.Trouble) {
	n.delegate.Support(trouble)
}

func (n *NoSupport) GetDelegate() *SupportDelegate {
	return n.delegate
}

func (n *NoSupport) String() string {
	return fmt.Sprintf("[%s@NoSupport]", n.name)
}
