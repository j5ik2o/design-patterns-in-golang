package strategy

import (
	"fmt"
	"testing"
)

func TestStrategy(t *testing.T) {
	winningStrategy := NewWinningStrategy()
	probeStrategy := NewProbeStrategy()

	player1 := NewPlayer("Taro", winningStrategy)
	player2 := NewPlayer("Hana", probeStrategy)

	for i := 0; i < 10000; i++ {
		nextHand1 := player1.NextHand().MustGet()
		nextHand2 := player2.NextHand().MustGet()
		if nextHand1.IsStrongerThan(*nextHand2) {
			fmt.Printf("Winner:%s\n", player1.String())
			player1.Win()
			player2.Lose()
		} else if nextHand1.IsStrongerThan(*nextHand1) {
			fmt.Printf("Winner:%s\n", player2.String())
			player1.Lose()
			player2.Win()
		} else {
			player1.Even()
			player2.Even()
		}
	}

	fmt.Println("Total result:")
	fmt.Println(player1.String())
	fmt.Println(player2.String())

}
