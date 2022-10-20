package flyweight

import (
	"fmt"
	"testing"
)

func TestFlyweight(t *testing.T) {
	bs := NewBigString("1928374650564738291")
	fmt.Println(bs.ToString())
}
