package framework

type UseF func()
type StringF func() string
type Product struct {
	Use    UseF
	String StringF
}

func NewProduct(use UseF, string StringF) *Product {
	return &Product{use, string}
}
