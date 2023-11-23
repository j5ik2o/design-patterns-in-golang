package delegate

import "testing"

func testPrint(p Print) {
	p.PrintWeak()
	p.PrintStrong()
}

func TestAdaptor(t *testing.T) {
	printBanner := NewPrintBanner(NewBanner("Hello"))
	testPrint(printBanner)
}
