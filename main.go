package main

import (
	"desgin-patterns-in-golang/template_method"
	"desgin-patterns-in-golang/template_method_old"
	"fmt"
)

func testDisplay(d template_method.Display) {
	d.Display()
}

func main() {
	fmt.Printf("Hello World\n")
	cd := template_method_old.NewCharDisplay('H')
	testDisplay(cd)
	sd := template_method_old.NewStringDisplay("def")
	testDisplay(sd)
}
