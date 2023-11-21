package delegate

import (
	"desgin-patterns-in-golang/chain_of_responsibility"
	"fmt"
)

type LimitSupport struct {
	name     string
	delegate *SupportDelegate
	limit    int
}

func NewLimitSupport(name string, limit int, next Support) *LimitSupport {
	var nextDelegate *SupportDelegate = nil
	if next != nil {
		nextDelegate = next.GetDelegate()
	}
	self := &LimitSupport{
		name:  name,
		limit: limit,
	}
	self.delegate = NewSupportDelegate(name, self.limitResolver, nextDelegate)
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

func (l *LimitSupport) GetDelegate() *SupportDelegate {
	return l.delegate
}

func (l *LimitSupport) String() string {
	return fmt.Sprintf("[%s@LimitSupport]", l.name)
}
