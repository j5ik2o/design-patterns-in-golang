package embbed

type Printer interface {
	open()
	print()
	close()
}
