package bridge

type Display interface {
	Open()
	Print()
	Close()
	Display()
}
