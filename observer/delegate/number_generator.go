package delegate

type NumberGenerator interface {
	AddObserver(observer Observer)
	DeleteObserver(observer Observer)
	NotifyObservers()
	GetNumber() uint
	Execute()
}
