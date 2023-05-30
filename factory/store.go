package factory

type PizzaStore struct {
	// Spf   SimplePizzaFactory
	Pizza Pizza
	IPizzaStore
}

/*
// PizzaStore
// Deprecated
func (ps *PizzaStore) PizzaStore(spf SimplePizzaFactory) {
	ps.Spf = spf
}
*/

func (ps *PizzaStore) OrderPizza(name string) Pizza {
	ps.Pizza = ps.CreatePizza(name)
	ps.Pizza.Prepare()
	ps.Pizza.Bake()
	ps.Pizza.Cut()
	ps.Pizza.Box()
	return ps.Pizza
}

type IPizzaStore interface {
	CreatePizza(string) Pizza
}
