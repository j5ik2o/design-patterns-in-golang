package main

import (
	"desgin-patterns-in-golang/template_method"
	"fmt"
)

func testDisplay(d template_method.Display) {
	d.Display()
}

func main() {
	fmt.Printf("Hello World\n")
	cd := template_method.NewCharDisplay('H')
	testDisplay(cd)
	sd := template_method.NewStringDisplay("def")
	testDisplay(sd)
}
