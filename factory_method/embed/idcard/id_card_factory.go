package idcard

import (
	"desgin-patterns-in-golang/factory_method/embed/framework"
	"fmt"
)

type IDCardFactory struct {
	*framework.Factory
}

func NewIDCardFactory() *IDCardFactory {
	self := &IDCardFactory{}
	self.Factory = framework.NewFactory(self.CreateProduct, self.RegisterProduct)
	return self
}

func (d *IDCardFactory) CreateProduct(owner string) *framework.Product {
	// NOTICE: IdCard型は捨てられる…。
	return NewIDCard(owner).Product
}

func (d *IDCardFactory) RegisterProduct(product *framework.Product) {
	fmt.Printf("%sを登録します。\n", product.String())
}
