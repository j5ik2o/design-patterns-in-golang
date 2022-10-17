package strategy

import "math/rand"

type Hand int

const (
	GUU Hand = iota
	CHO
	PAA
)

func (h *Hand) Name() string {
	var result string
	switch *h {
	case GUU:
		result = "グー"
	case CHO:
		result = "チョキ"
	case PAA:
		result = "パー"
	}
	return result
}

func GetHand(value int) Hand {
	var result Hand
	switch value {
	case 0:
		result = GUU
	case 1:
		result = CHO
	case 2:
		result = PAA
	default:
		panic("not found")
	}
	return result
}

func (h *Hand) HandValue() int {
	return (int)(*h)
}

func (h *Hand) Fight(other Hand) int {
	if *h == other {
		return 0
	} else if (h.HandValue()+1)%3 == h.HandValue() {
		return 1
	} else {
		return -1
	}
}

func (h *Hand) IsStrongerThan(other Hand) bool {
	return h.Fight(other) == 1
}

func (h *Hand) IsWeakerThan(other Hand) bool {
	return h.Fight(other) == -1
}

func (h *Hand) toString() string {
	return h.Name()
}

type Strategy interface {
	NextHand() Hand
}

type WinningStrategy struct {
	won      bool
	prevHand *Hand
}

func NewWinningStrategy() *WinningStrategy {
	return &WinningStrategy{
		won:      false,
		prevHand: nil,
	}
}

func (ws *WinningStrategy) NextHand() *Hand {
	if !ws.won {
		next := GetHand(rand.Int())
		ws.prevHand = &next
	}
	return ws.prevHand
}

func (ws *WinningStrategy) Study(win bool) {
	ws.won = win
}

type ProbeStrategy struct {
	prevHandValue    int
	currentHandValue int
	history          [][]int
}
