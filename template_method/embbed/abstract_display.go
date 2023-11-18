package embbed

type AbstractDisplay struct {
	printer Printer
}

func (o *AbstractDisplay) Display() {
	o.printer.open()
	for i := 0; i < 5; i++ {
		o.printer.print()
	}
	o.printer.close()
}
