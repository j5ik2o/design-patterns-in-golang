package delegate

import (
	"github.com/samber/mo"
	"math/rand"
)

type WinningStrategy struct {
	won      bool
	prevHand mo.Option[*Hand]
}

func NewWinningStrategy() *WinningStrategy {
	return &WinningStrategy{
		won:      false,
		prevHand: mo.None[*Hand](),
	}
}

func (ws *WinningStrategy) NextHand() mo.Option[*Hand] {
	if !ws.won {
		result := GetHand(rand.Intn(2))
		ws.prevHand = mo.Some(&result)
	}
	return ws.prevHand
}

func (ws *WinningStrategy) Study(win bool) {
	ws.won = win
}
