package bridge

import (
	"math/rand"
)

type RandomCountDisplay struct {
	*CountDisplay
}

func NewRandomCountDisplay(underlying DisplayDelegate) *RandomCountDisplay {
	return &RandomCountDisplay{NewCountDisplay(underlying)}
}

func (r *RandomCountDisplay) RandomDisplay(times int) {
	n := rand.Intn(times)
	r.MultiDisplay(n)
}
