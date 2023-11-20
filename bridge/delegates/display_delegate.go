package delegates

type DisplayDelegate interface {
	RawOpen()
	RawPrint()
	RawClose()
}
