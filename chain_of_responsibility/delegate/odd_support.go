package delegate

import (
	"desgin-patterns-in-golang/chain_of_responsibility"
	"fmt"
)

type OddSupport struct {
	name     string
	delegate *SupportDelegate
}

func NewOddSupport(name string, next Support) *OddSupport {
	var nextDelegate *SupportDelegate = nil
	if next != nil {
		nextDelegate = next.GetDelegate()
	}
	return &OddSupport{
		name,
		NewSupportDelegate(name, oddResolver, nextDelegate),
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

func (o *OddSupport) GetDelegate() *SupportDelegate {
	return o.delegate
}

func (o *OddSupport) String() string {
	return fmt.Sprintf("[%s@OddSupport]", o.name)
}
