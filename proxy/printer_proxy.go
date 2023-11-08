package proxy

type PrinterProxy struct {
	name       string
	underlying *Printer
}

func NewPrinterProxy(name string) *PrinterProxy {
	return &PrinterProxy{name, nil}
}

func (p *PrinterProxy) Realize() {
	if p.underlying == nil {
		p.underlying = NewPrinter(p.name)
	}
}

func (p *PrinterProxy) SetPrinterName(name string) {
	if p.underlying != nil {
		p.underlying.SetPrinterName(name)
	}
	p.name = name
}

func (p *PrinterProxy) GetPrinterName() string {
	return p.name
}

func (p *PrinterProxy) Print(text string) {
	p.Realize()
	p.underlying.Print(text)
}
