package supports

import (
	"desgin-patterns-in-golang/chain_of_responsibility"
	delegate2 "desgin-patterns-in-golang/chain_of_responsibility/delegate"
	"fmt"
)

type OddSupport struct {
	name     string
	delegate *delegate2.SupportDelegate
}

func NewOddSupport(name string, next delegate2.Support) *OddSupport {
	var nextDelegate *delegate2.SupportDelegate = nil
	if next != nil {
		nextDelegate = next.GetDelegate()
	}
	return &OddSupport{
		name,
		delegate2.NewSupportDelegate(name, oddResolver, nextDelegate),
	}
}

func oddResolver(t *chain_of_responsibility.Trouble) bool {
	if t.GetNumber()%2 == 1 {
		return true
	}
	return false
}

func (o *OddSupport) Support(trouble *chain_of_responsibility.Trouble) {
	o.delegate.Support(trouble)
}

func (o *OddSupport) GetDelegate() *delegate2.SupportDelegate {
	return o.delegate
}

func (o *OddSupport) String() string {
	return fmt.Sprintf("[%s@OddSupport]", o.name)
}
