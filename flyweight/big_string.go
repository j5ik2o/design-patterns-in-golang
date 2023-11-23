package flyweight

import "strings"

var bigCharFactorySingleton *BigCharFactory = NewBigCharFactory()

type BigString struct {
	bigChars []*BigChar
}

func NewBigString(string string) *BigString {
	bigString := &BigString{}
	slice := strings.Split(string, "")
	for i := 0; i < len(slice); i++ {
		bigString.bigChars = append(bigString.bigChars, bigCharFactorySingleton.GetBigChar(slice[i]))
	}
	return bigString
}

func (bs *BigString) ToString() string {
	var result string
	for i := 0; i < len(bs.bigChars); i++ {
		result += bs.bigChars[i].String()
	}
	return result
}
