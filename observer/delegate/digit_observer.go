package delegate

import (
	"fmt"
	"time"
)

type DigitObserver struct {
}

func (d DigitObserver) Update(generator NumberGenerator) {
	fmt.Printf("DigitObserver:%d\n", generator.GetNumber())
	time.Sleep(100 * time.Millisecond)
}
