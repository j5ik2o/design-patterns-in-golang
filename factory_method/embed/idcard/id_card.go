package idcard

import (
	"desgin-patterns-in-golang/factory_method/embed/framework"
	"fmt"
)

type IDCard struct {
	*framework.Product
	owner string
}

func NewIDCard(owner string) *IDCard {
	fmt.Println(owner + "のカードを作ります。")
	self := &IDCard{owner: owner}
	self.Product = framework.NewProduct(self.Use, self.String)
	return self
}

func (c *IDCard) Use() {
	fmt.Println(c.owner + "のカードを使います。")
}

func (c *IDCard) GetOwner() string {
	return c.owner
}

func (c *IDCard) String() string {
	return fmt.Sprintf("IDCard{ owner = %s }", c.owner)
}
