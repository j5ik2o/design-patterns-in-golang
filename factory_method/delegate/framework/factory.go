package framework

type Factory struct {
	delegate FactoryDelegate
}

func NewFactory(delegate FactoryDelegate) *Factory {
	return &Factory{delegate}
}

func (f *Factory) Create(owner string) Product {
	p := f.delegate.CreateProduct(owner)
	f.delegate = f.delegate.RegisterProduct(p)
	return p
}
