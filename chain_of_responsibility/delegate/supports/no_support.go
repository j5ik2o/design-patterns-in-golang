package supports

import (
	"desgin-patterns-in-golang/chain_of_responsibility"
	delegate2 "desgin-patterns-in-golang/chain_of_responsibility/delegate"
	"fmt"
)

type NoSupport struct {
	name     string
	delegate *delegate2.SupportDelegate
}

func NewNoSupport(name string, next delegate2.Support) *NoSupport {
	var nextDelegate *delegate2.SupportDelegate = nil
	if next != nil {
		nextDelegate = next.GetDelegate()
	}
	return &NoSupport{
		name,
		delegate2.NewSupportDelegate(name, resolver, nextDelegate),
	}
}

func resolver(t *chain_of_responsibility.Trouble) bool {
	return false
}

func (n *NoSupport) Support(trouble *chain_of_responsibility.Trouble) {
	n.delegate.Support(trouble)
}

func (n *NoSupport) GetDelegate() *delegate2.SupportDelegate {
	return n.delegate
}

func (n *NoSupport) String() string {
	return fmt.Sprintf("[%s@NoSupport]", n.name)
}
