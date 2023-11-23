package delegate

import "github.com/samber/mo"

type Strategy interface {
	NextHand() mo.Option[*Hand]
	Study(win bool)
}
