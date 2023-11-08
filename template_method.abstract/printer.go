package template_method_abstract

type Printer interface {
	open()
	print()
	close()
}
