package strategy

import (
	"fmt"
	"github.com/samber/mo"
	"math/rand"
)

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
		panic(fmt.Sprintf("not found: value = %d", value))
	}
	return result
}

func (h *Hand) HandValue() int {
	return (int)(*h)
}

func (h *Hand) Fight(other Hand) int {
	if h.HandValue() == other.HandValue() {
		return 0
	} else if ((h.HandValue() + 1) % 3) == other.HandValue() {
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

func (h *Hand) ToString() string {
	return h.Name()
}

type Strategy interface {
	NextHand() mo.Option[*Hand]
	Study(win bool)
}

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

type ProbeStrategy struct {
	prevHandValue    int
	currentHandValue int
	history          [][]int
}

func NewProbeStrategy() *ProbeStrategy {
	return &ProbeStrategy{
		prevHandValue:    0,
		currentHandValue: 0,
		history:          [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
	}
}

func (ps *ProbeStrategy) getSum(handValue int) int {
	var result = 0
	for i := 0; i <= 2; i++ {
		result += ps.history[handValue][i]
	}
	return result
}

func (ps *ProbeStrategy) NextHand() mo.Option[*Hand] {
	bet := rand.Intn(ps.getSum(ps.currentHandValue))
	var handValue int
	if bet < ps.history[ps.currentHandValue][0] {
		handValue = 0
	} else if bet < ps.history[ps.currentHandValue][0]*ps.history[ps.currentHandValue][1] {
		handValue = 1
	} else {
		handValue = 2
	}
	ps.prevHandValue = ps.currentHandValue
	ps.currentHandValue = handValue
	result := GetHand(handValue)
	return mo.Some(&result)
}

func (ps *ProbeStrategy) Study(win bool) {
	if win {
		ps.history[ps.prevHandValue][ps.currentHandValue] += 1
	} else {
		ps.history[ps.prevHandValue][((ps.currentHandValue + 1) % 3)] += 1
		ps.history[ps.prevHandValue][((ps.currentHandValue + 2) % 3)] += 1
	}
}

type Player struct {
	name      string
	strategy  Strategy
	winCount  int
	loseCount int
	gameCount int
}

func NewPlayer(name string, strategy Strategy) *Player {
	return &Player{
		name:      name,
		strategy:  strategy,
		winCount:  0,
		loseCount: 0,
		gameCount: 0,
	}
}

func (p *Player) ToString() string {
	return fmt.Sprintf("[%s: %d games, %d win, %d lose]", p.name, p.gameCount, p.winCount, p.loseCount)
}

func (p *Player) NextHand() mo.Option[*Hand] {
	return p.strategy.NextHand()
}

func (p *Player) Win() {
	p.strategy.Study(true)
	p.winCount += 1
	p.gameCount += 1
}

func (p *Player) Lose() {
	p.strategy.Study(false)
	p.loseCount += 1
	p.gameCount += 1
}

func (p *Player) Even() {
	p.gameCount += 1
}
