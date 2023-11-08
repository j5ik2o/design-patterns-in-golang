package chain_of_responsibility

import (
	"fmt"
	"github.com/samber/mo"
)

type OddSupport struct {
	name string
	next mo.Option[Support]
}

func NewOddSupport(name string, next mo.Option[Support]) *OddSupport {
	return &OddSupport{
		name,
		next,
	}
}

func (o *OddSupport) String() string {
	return fmt.Sprintf("[%s@OddSupport]", o.name)
}

func (o *OddSupport) Done(trouble *Trouble) {
	fmt.Printf("%s is resolved by %s.\n", trouble, o.String())
}

func (o *OddSupport) Fail(trouble *Trouble) {
	fmt.Printf("%s cannot be resolved.\n", trouble)
}

func (o *OddSupport) Next() mo.Option[Support] {
	return o.next
}

func (o *OddSupport) Resolve(trouble *Trouble) bool {
	if trouble.GetNumber()%2 == 1 {
		return true
	}
	return false
}

func (o *OddSupport) Support(trouble *Trouble) {
	support(o, trouble)
}
