package chain_of_responsibility

import (
	"fmt"
	"github.com/samber/mo"
)

type LimitSupport struct {
	name  string
	next  mo.Option[Support]
	limit int
}

func NewLimitSupport(name string, limit int, next mo.Option[Support]) *LimitSupport {
	return &LimitSupport{
		name,
		next,
		limit,
	}
}

func (l *LimitSupport) String() string {
	return fmt.Sprintf("[%s@LimitSupport]", l.name)
}

func (l *LimitSupport) Done(trouble *Trouble) {
	fmt.Printf("%s is resolved by %s.\n", trouble, l.String())
}

func (l *LimitSupport) Fail(trouble *Trouble) {
	fmt.Printf("%s cannot be resolved.\n", trouble)
}

func (l *LimitSupport) Next() mo.Option[Support] {
	return l.next
}

func (l *LimitSupport) Resolve(trouble *Trouble) bool {
	if trouble.GetNumber() < l.limit {
		return true
	}
	return false
}

func (l *LimitSupport) Support(trouble *Trouble) {
	support(l, trouble)
}
