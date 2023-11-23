package delegate

type Colleague interface {
	GetName() string
	OnChanged(msg string)
	Run()
}
