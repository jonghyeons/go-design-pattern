package decorator

type Espresso struct {
	Description string
	Beverage
}

// NewEspresso
// Constructor 대응
func NewEspresso() Beverage {
	e := &Espresso{
		Description: "에스프레소",
	}
	return e
}

func (e *Espresso) Cost() float64 {
	return 1.99
}

func (e *Espresso) GetDescription() string {
	return e.Description
}
