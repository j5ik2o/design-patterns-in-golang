package delegate

import (
	"fmt"
	"github.com/samber/mo"
)

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

func (p *Player) String() string {
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
