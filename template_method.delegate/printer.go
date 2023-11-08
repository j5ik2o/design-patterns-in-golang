package template_method_delegate

type Printer interface {
	open()
	print()
	close()
}
