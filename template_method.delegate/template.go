package template_method_delegate

type Template struct {
	printer Printer
}

func NewTemplate(printer Printer) *Template {
	return &Template{printer: printer}
}

func (o *Template) display() {
	o.printer.open()
	for i := 0; i < 5; i++ {
		o.printer.print()
	}
	o.printer.close()
}
