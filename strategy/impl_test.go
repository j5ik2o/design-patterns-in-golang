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
		nextHand1 := player1.NextHand()
		nextHand2 := player2.NextHand()
		if nextHand1.IsStrongerThan(*nextHand2) {
			fmt.Printf("Winner:%s\n", player1.ToString())
			player1.Win()
			player2.Lose()
		} else if nextHand1.IsStrongerThan(*nextHand1) {
			fmt.Printf("Winner:%s\n", player2.ToString())
			player1.Lose()
			player2.Win()
		} else {
			player1.Even()
			player2.Even()
		}
	}

	fmt.Println("Total result:")
	fmt.Println(player1.ToString())
	fmt.Println(player2.ToString())

}
