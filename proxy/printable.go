package proxy

type Printable interface {
	SetPrinterName(name string)
	GetPrinterName() string
	Print(text string)
}
