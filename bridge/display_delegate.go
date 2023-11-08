package bridge

type DisplayDelegate interface {
	RawOpen()
	RawPrint()
	RawClose()
}
