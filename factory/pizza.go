package factory

type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
}

var (
	_ Pizza = (*CheesePizza)(nil)
	_ Pizza = (*PepperoniPizza)(nil)
	_ Pizza = (*ClamPizza)(nil)
	_ Pizza = (*VeggiePizza)(nil)
)

type CheesePizza struct{}

func (c CheesePizza) Prepare() {}
func (c CheesePizza) Bake()    {}
func (c CheesePizza) Cut()     {}
func (c CheesePizza) Box()     {}

type PepperoniPizza struct{}

func (p PepperoniPizza) Prepare() {}
func (p PepperoniPizza) Bake()    {}
func (p PepperoniPizza) Cut()     {}
func (p PepperoniPizza) Box()     {}

type ClamPizza struct{}

func (c ClamPizza) Prepare() {}
func (c ClamPizza) Bake()    {}
func (c ClamPizza) Cut()     {}
func (c ClamPizza) Box()     {}

type VeggiePizza struct{}

func (v VeggiePizza) Prepare() {}
func (v VeggiePizza) Bake()    {}
func (v VeggiePizza) Cut()     {}
func (v VeggiePizza) Box()     {}

var (
	_ Pizza = (*NYStyleCheesePizza)(nil)
	_ Pizza = (*NYStyleVeggiePizza)(nil)
	_ Pizza = (*NYStyleClamPizza)(nil)
	_ Pizza = (*NYStylePepperoniPizza)(nil)
)

type NYStyleCheesePizza struct{}

func (N NYStyleCheesePizza) Prepare() {}
func (N NYStyleCheesePizza) Bake()    {}
func (N NYStyleCheesePizza) Cut()     {}
func (N NYStyleCheesePizza) Box()     {}

type NYStyleVeggiePizza struct{}

func (N NYStyleVeggiePizza) Prepare() {}
func (N NYStyleVeggiePizza) Bake()    {}
func (N NYStyleVeggiePizza) Cut()     {}
func (N NYStyleVeggiePizza) Box()     {}

type NYStyleClamPizza struct{}

func (N NYStyleClamPizza) Prepare() {}
func (N NYStyleClamPizza) Bake()    {}
func (N NYStyleClamPizza) Cut()     {}
func (N NYStyleClamPizza) Box()     {}

type NYStylePepperoniPizza struct{}

func (N NYStylePepperoniPizza) Prepare() {}
func (N NYStylePepperoniPizza) Bake()    {}
func (N NYStylePepperoniPizza) Cut()     {}
func (N NYStylePepperoniPizza) Box()     {}
