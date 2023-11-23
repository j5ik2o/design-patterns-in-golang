package flyweight

import (
	"bufio"
	"fmt"
	"os"
)

type BigChar struct {
	charName string
	fontData string
}

func (bc *BigChar) String() string {
	return bc.fontData
}

func readFontData(charName string) string {
	// fmt.Printf("read font data: %s\n", charName)
	fileName := fmt.Sprintf("big%s.txt", charName)
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	var result string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result += scanner.Text() + "\n"
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
