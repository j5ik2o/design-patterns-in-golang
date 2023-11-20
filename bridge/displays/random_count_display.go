package displays

import (
	"desgin-patterns-in-golang/bridge/delegates"
	"math/rand"
)

type RandomCountDisplay struct {
	*CountDisplay
}

func NewRandomCountDisplay(underlying delegates.DisplayDelegate) *RandomCountDisplay {
	return &RandomCountDisplay{NewCountDisplay(underlying)}
}

func (r *RandomCountDisplay) RandomDisplay(times int) {
	n := rand.Intn(times)
	r.MultiDisplay(n)
}
