package templateMethod

import "fmt"

type Coffee struct {
	CaffeineBeverage
}

var _ CaffeineBeverage = (*Coffee)(nil)

func (c *Coffee) PrepareRecipe() {
	c.BoilWater()
	c.Brew()
	c.PourInCup()
	c.AddCondiments()
}

func (c *Coffee) BoilWater() {
	fmt.Println("물 끓이는 중")
}

func (c *Coffee) Brew() {
	fmt.Println("필터로 커피를 우려내는 중")
}

func (c *Coffee) PourInCup() {
	fmt.Println("컵에 따르는 중")
}

func (c *Coffee) AddCondiments() {
	fmt.Println("설탕과 우유를 추가하는 중")
}
