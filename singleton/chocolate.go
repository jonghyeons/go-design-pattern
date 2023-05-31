package singleton

type ChocolateBoiler struct {
	Empty  bool
	Boiled bool
}

func (c *ChocolateBoiler) Constructor() {
	c.Empty = true
	c.Boiled = true
}

func (c *ChocolateBoiler) Fill() {
	if c.IsEmpty() {
		c.Empty = false
		c.Boiled = false
	}
}

func (c *ChocolateBoiler) Drain() {
	if c.IsEmpty() && c.IsBoiled() {
		c.Empty = true
	}
}

func (c *ChocolateBoiler) Boil() {
	if !c.IsEmpty() && !c.IsBoiled() {
		c.Empty = false
	}
}

func (c *ChocolateBoiler) IsEmpty() bool {
	return c.Empty
}

func (c *ChocolateBoiler) IsBoiled() bool {
	return c.Boiled
}

var chocolateBoiler *ChocolateBoiler

type ChocolateBoilerSingleton struct{}

func (cbs *ChocolateBoilerSingleton) GetInstance() *ChocolateBoiler {
	if chocolateBoiler == nil {
		return new(ChocolateBoiler)
	}
	return chocolateBoiler
}
