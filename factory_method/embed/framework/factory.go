package framework

type CreateProductF func(owner string) *Product
type RegisterProductF func(product *Product)

type Factory struct {
	CreateProduct   CreateProductF
	RegisterProduct RegisterProductF
}

func NewFactory(createProduct CreateProductF, registerProduct RegisterProductF) *Factory {
	return &Factory{createProduct, registerProduct}
}

func (f *Factory) Create(owner string) *Product {
	p := f.CreateProduct(owner)
	f.RegisterProduct(p)
	return p
}
