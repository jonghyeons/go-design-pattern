package factory

type SimplePizzaFactory struct {
	Pizza Pizza
}

func (spf *SimplePizzaFactory) CreatePizza(name string) Pizza {
	if name == "cheese" {
		return new(CheesePizza)
	} else if name == "pepperoni" {
		return new(PepperoniPizza)
	} else if name == "clam" {
		return new(ClamPizza)
	} else if name == "veggie" {
		return new(VeggiePizza)
	}
	return spf.Pizza
}

type ISimplePizzaFactory interface {
	OrderPizza(string)
}
