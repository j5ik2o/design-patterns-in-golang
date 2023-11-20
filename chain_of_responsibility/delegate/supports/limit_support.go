package supports

import (
	"desgin-patterns-in-golang/chain_of_responsibility"
	delegate2 "desgin-patterns-in-golang/chain_of_responsibility/delegate"
	"fmt"
)

type LimitSupport struct {
	name     string
	delegate *delegate2.SupportDelegate
	limit    int
}

func NewLimitSupport(name string, limit int, next delegate2.Support) *LimitSupport {
	var nextDelegate *delegate2.SupportDelegate = nil
	if next != nil {
		nextDelegate = next.GetDelegate()
	}
	self := &LimitSupport{
		name:  name,
		limit: limit,
	}
	self.delegate = delegate2.NewSupportDelegate(name, self.limitResolver, nextDelegate)
	return self
}

func (l *LimitSupport) limitResolver(t *chain_of_responsibility.Trouble) bool {
	if t.GetNumber() < l.limit {
		return true
	}
	return false
}

func (l *LimitSupport) Support(trouble *chain_of_responsibility.Trouble) {
	l.delegate.Support(trouble)
}

func (l *LimitSupport) GetDelegate() *delegate2.SupportDelegate {
	return l.delegate
}

func (l *LimitSupport) String() string {
	return fmt.Sprintf("[%s@LimitSupport]", l.name)
}
