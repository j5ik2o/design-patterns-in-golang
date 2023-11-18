package idcard

import "fmt"

type IDCard struct {
	owner string
}

func NewIDCard(owner string) *IDCard {
	fmt.Println(owner + "のカードを作ります。")
	return &IDCard{
		owner: owner,
	}
}

func (c *IDCard) Use() {
	fmt.Println(c.owner + "のカードを使います。")
}

func (c *IDCard) GetOwner() string {
	return c.owner
}

func (c *IDCard) ToString() string {
	return "[IDCard:" + c.owner + "]"
}
