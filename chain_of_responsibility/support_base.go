package chain_of_responsibility

import "github.com/samber/mo"

type SupportBase interface {
	Done(trouble *Trouble)
	Fail(trouble *Trouble)
	Next() mo.Option[Support]
}
