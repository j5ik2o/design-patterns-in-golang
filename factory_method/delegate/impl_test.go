package delegate

import (
	"desgin-patterns-in-golang/factory_method/delegate/framework"
	"desgin-patterns-in-golang/factory_method/delegate/idcard"
	"testing"
)

func TestFactoryMethod(t *testing.T) {
	factory := framework.NewFactory(idcard.NewIDCardFactoryDelegate())
	card1 := factory.Create("結城浩")
	card2 := factory.Create("とむら")
	card3 := factory.Create("佐藤花子")
	card1.Use()
	card2.Use()
	card3.Use()
}
