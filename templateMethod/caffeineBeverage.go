package templateMethod

type CaffeineBeverage interface {
	PrepareRecipe()
	BoilWater()
	PourInCup()
	Brew()
	AddCondiments()
}
