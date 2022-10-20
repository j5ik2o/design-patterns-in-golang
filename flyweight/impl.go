package flyweight

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type BigChar struct {
	charName string
	fontData string
}

func (bc *BigChar) ToString() string {
	return bc.fontData
}

func readFontData(charName string) string {
	fmt.Printf("read font data: %s", charName)
	fileName := fmt.Sprintf("flyweight/big%s.txt", charName)
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var result string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result += scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return result
}

func NewBigChar(charName string) *BigChar {
	return &BigChar{
		charName: charName,
		fontData: readFontData(charName),
	}
}

type BigCharFactory struct {
	pool map[string]*BigChar
}

func NewBigCharFactory() *BigCharFactory {
	return &BigCharFactory{
		pool: make(map[string]*BigChar),
	}
}

func (f *BigCharFactory) GetBigChar(charName string) *BigChar {
	bc, ok := f.pool[charName]
	if !ok {
		bc = NewBigChar(charName)
		f.pool[charName] = bc
	}
	return bc
}

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
		result += bs.bigChars[i].ToString()
	}
	return result
}
