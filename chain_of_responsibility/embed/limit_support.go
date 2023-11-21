package embed

import (
	"desgin-patterns-in-golang/chain_of_responsibility"
)

type LimitSupport struct {
	*SupportBase
	limit int
}

func NewLimitSupport(name string, limit int, next *SupportBase) *LimitSupport {
	self := &LimitSupport{
		limit: limit,
	}
	self.SupportBase = NewSupportBase(name, self.resolve, next)
	return self
}

func (l *LimitSupport) resolve(trouble *chain_of_responsibility.Trouble) bool {
	return trouble.GetNumber() < l.limit
}
