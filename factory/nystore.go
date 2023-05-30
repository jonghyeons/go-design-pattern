package factory

type NYPizzaStore struct {
	PizzaStore
}

func (ny *NYPizzaStore) CreatePizza(name string) Pizza {
	if name == "cheese" {
		return new(NYStyleCheesePizza)
	} else if name == "pepperoni" {
		return new(NYStylePepperoniPizza)
	} else if name == "clam" {
		return new(NYStyleClamPizza)
	} else if name == "veggie" {
		return new(NYStyleVeggiePizza)
	}

	pizza := new(Pizza)
	return *pizza
}
