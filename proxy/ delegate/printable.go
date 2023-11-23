package _delegate

type Printable interface {
	SetPrinterName(name string)
	GetPrinterName() string
	Print(text string)
}
