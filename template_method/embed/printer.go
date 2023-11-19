package embed

type Printer interface {
	open()
	print()
	close()
}
