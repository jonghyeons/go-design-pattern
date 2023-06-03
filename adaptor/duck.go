package adaptor

import "fmt"

type Duck interface {
	Quack()
	Fly()
}

type MallardDuck struct {
	Duck
}

var _ Duck = (*MallardDuck)(nil)

func (m MallardDuck) Quack() {
	fmt.Println("꽥")
}

func (m MallardDuck) Fly() {
	fmt.Println("날고 있어요")
}
