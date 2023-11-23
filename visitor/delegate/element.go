package delegate

type Element interface {
	Accept(visitor Visitor)
}
