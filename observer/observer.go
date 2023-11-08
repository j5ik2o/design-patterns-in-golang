package observer

type Observer interface {
	Update(generator NumberGenerator)
}
