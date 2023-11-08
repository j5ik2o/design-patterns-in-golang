package chain_of_responsibility

import (
	"fmt"
	"github.com/samber/mo"
)

type NoSupport struct {
	name string
	next mo.Option[Support]
}

func NewNoSupport(name string, next mo.Option[Support]) *NoSupport {
	return &NoSupport{
		name,
		next,
	}
}

func (n *NoSupport) String() string {
	return fmt.Sprintf("[%s@NoSupport]", n.name)
}

func (n *NoSupport) Done(trouble *Trouble) {
	fmt.Printf("%s is resolved by %s.\n", trouble, n.String())
}

func (n *NoSupport) Fail(trouble *Trouble) {
	fmt.Printf("%s cannot be resolved.\n", trouble)
}

func (n *NoSupport) Next() mo.Option[Support] {
	return n.next
}

func (n *NoSupport) Resolve(trouble *Trouble) bool {
	return false
}

func (n *NoSupport) Support(trouble *Trouble) {
	support(n, trouble)
}

func support(s Support, trouble *Trouble) {
	if s.Resolve(trouble) {
		s.Done(trouble)
	} else if s.Next().IsPresent() {
		s.Next().MustGet().Support(trouble)
	} else {
		s.Fail(trouble)
	}
}
