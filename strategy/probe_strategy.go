package strategy

import (
	"github.com/samber/mo"
	"math/rand"
)

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
