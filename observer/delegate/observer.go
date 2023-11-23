package delegate

type Observer interface {
	Update(generator NumberGenerator)
}
