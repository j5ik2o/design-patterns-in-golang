package template_method_embbed

type Printer interface {
	open()
	print()
	close()
}
