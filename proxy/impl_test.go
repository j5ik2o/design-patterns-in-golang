package proxy

import (
	"fmt"
	"testing"
)

func process(p Printable) {
	fmt.Printf("名前は現在%sです。\n", p.GetPrinterName())
	p.SetPrinterName("Bob")
	fmt.Printf("名前は現在%sです。\n", p.GetPrinterName())
	p.Print("Hello, world.")
}

func TestProxy(t *testing.T) {
	proxy := NewPrinterProxy("Alice")
	process(proxy)
}
