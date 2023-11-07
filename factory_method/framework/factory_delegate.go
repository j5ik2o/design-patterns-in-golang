package framework

type FactoryDelegate interface {
	CreateProduct(owner string) Product
	RegisterProduct(product Product)
}
