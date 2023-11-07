package idcard

import (
	"desgin-patterns-in-golang/factory_method/framework"
	"fmt"
)

type IDCardFactoryDelegate struct {
}

func NewIDCardFactoryDelegate() *IDCardFactoryDelegate {
	return &IDCardFactoryDelegate{}
}

func (d *IDCardFactoryDelegate) CreateProduct(owner string) framework.Product {
	return NewIDCard(owner)
}

func (d *IDCardFactoryDelegate) RegisterProduct(product framework.Product) framework.FactoryDelegate {
	fmt.Println(product.ToString() + "を登録します。")
	return d
}
