package delegate

import (
	"fmt"
	"time"
)

type GraphObserver struct {
}

func (g GraphObserver) Update(generator NumberGenerator) {
	fmt.Print("GraphObserver:")
	count := generator.GetNumber()
	for i := uint(0); i < count; i++ {
		fmt.Print("*")
	}
	fmt.Println("")
	time.Sleep(100 * time.Millisecond)
}
