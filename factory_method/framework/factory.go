package framework

type Factory struct {
	delegate FactoryDelegate
}

func NewFactory(delegate FactoryDelegate) *Factory {
	return &Factory{
		delegate: delegate,
	}
}

func (f *Factory) Create(owner string) Product {
	p := f.delegate.CreateProduct(owner)
	f.delegate.RegisterProduct(p)
	return p
}
