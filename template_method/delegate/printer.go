package delegate

type Printer interface {
	open()
	print()
	close()
}
