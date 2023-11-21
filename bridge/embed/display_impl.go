package embed

type DisplayImpl interface {
	RawOpen()
	RawPrint()
	RawClose()
}
