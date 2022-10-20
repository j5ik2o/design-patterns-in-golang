package proxy

import (
	"fmt"
	"time"
)

type Printable interface {
	SetPrinterName(name string)
	GetPrinterName() string
	Print(text string)
}

type Printer struct {
	name string
}

func NewPrinter(name string) *Printer {
	heavyJob(fmt.Sprintf("Printerのインスタンス(%s)を生成中", name))
	return &Printer{name}
}

func heavyJob(msg string) {
	fmt.Printf("%s", msg)
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Print(".")
	}
	fmt.Println("完了。")
}

func (p *Printer) SetPrinterName(name string) {
	p.name = name
}

func (p *Printer) GetPrinterName() string {
	return p.name
}

func (p *Printer) Print(text string) {
	fmt.Printf("=== %s ===\n", p.name)
	fmt.Println(text)
}

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
