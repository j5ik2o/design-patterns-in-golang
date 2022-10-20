package factory_method

import "fmt"

type Product interface {
	Use()
	ToString() string
}

type FactoryBase interface {
	CreateProduct(owner string) Product
	RegisterProduct(product Product)
}

type Factory interface {
	FactoryBase
	Create(owner string) Product
}

type IdCard struct {
	owner string
}

func NewIdCard(owner string) *IdCard {
	fmt.Printf("%sのカードを作ります。\n", owner)
	return &IdCard{owner}
}

func (i *IdCard) GetOwner() string {
	return i.owner
}

func (i *IdCard) ToString() string {
	return fmt.Sprintf("[IdCard:%s]", i.owner)
}

func (i *IdCard) Use() {
	fmt.Printf("%sを使います。\n", i.owner)
}

type IdCardFactory struct {
}

func (i *IdCardFactory) Create(owner string) Product {
	p := i.CreateProduct(owner)
	i.RegisterProduct(p)
	return p
}

func (i *IdCardFactory) CreateProduct(owner string) Product {
	return NewIdCard(owner)
}

func (i *IdCardFactory) RegisterProduct(product Product) {
	fmt.Printf("%sを登録しました。\n", product.ToString())
}
