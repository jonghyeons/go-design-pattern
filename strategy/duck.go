package strategy

type Duck struct {
	FlyBehavior
	QuackBehavior
}

func (d *Duck) PerformQuack() {
	d.Quack()
}

func (d *Duck) PerformFly() {
	d.Fly()
}

func (d *Duck) Swim()    {}
func (d *Duck) Display() {}
