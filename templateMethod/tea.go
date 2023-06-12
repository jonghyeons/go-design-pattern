package templateMethod

import "fmt"

type Tea struct {
	CaffeineBeverage
}

var _ CaffeineBeverage = (*Tea)(nil)

func (t *Tea) PrepareRecipe() {
	t.BoilWater()
	t.Brew()
	t.PourInCup()
	t.AddCondiments()
}

func (t *Tea) BoilWater() {
	fmt.Println("물 끓이는 중")
}

func (t *Tea) Brew() {
	fmt.Println("찻잎을 우려내는 중")
}

func (t *Tea) PourInCup() {
	fmt.Println("컵에 따르는 중")
}

func (t *Tea) AddCondiments() {
	fmt.Println("레몬을 추가하는 중")
}
