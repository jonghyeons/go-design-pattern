package decorator

type Mocha struct {
	Beverage
}

// NewMocha
// Constructor 대응
func NewMocha(b Beverage) Beverage {
	m := &Mocha{b}
	return m
}

func (m *Mocha) GetDescription() string {
	return m.Beverage.GetDescription() + ", 모카"
}

func (m *Mocha) Cost() float64 {
	return m.Beverage.Cost() + .20
}
