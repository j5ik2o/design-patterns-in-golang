package idcard

import (
	framework2 "desgin-patterns-in-golang/factory_method/delegate/framework"
	"fmt"
)

type IDCardFactoryDelegate struct {
}

func NewIDCardFactoryDelegate() *IDCardFactoryDelegate {
	return &IDCardFactoryDelegate{}
}

func (d *IDCardFactoryDelegate) CreateProduct(owner string) framework2.Product {
	return NewIDCard(owner)
}

func (d *IDCardFactoryDelegate) RegisterProduct(product framework2.Product) framework2.FactoryDelegate {
	fmt.Println(product.ToString() + "を登録します。")
	return d
}
