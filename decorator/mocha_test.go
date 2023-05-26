package decorator

import (
	"fmt"
	"testing"
)

func TestMocha(t *testing.T) {
	beverage := NewEspresso()
	fmt.Println(beverage.GetDescription(), "$", beverage.Cost())

	var beverage2 Beverage
	beverage2 = NewHouseBlend()
	fmt.Println(beverage2.GetDescription(), "$", beverage2.Cost())

	var beverage3 Beverage
	beverage3 = NewEspresso()
	beverage3 = NewMocha(beverage3)
	beverage3 = NewMocha(beverage3)
	fmt.Println(beverage3.GetDescription(), "$", beverage3.Cost())
}
