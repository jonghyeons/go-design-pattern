package strategy

import "fmt"

type FlyBehavior interface {
	Fly()
}

type FlyWithWings struct {
	FlyBehavior
}

func (fwn *FlyWithWings) Fly() {
	fmt.Println("날아요")
}

type FlyNoWay struct {
	FlyBehavior
}

func (fnw *FlyNoWay) Fly() {
	fmt.Println("못날아요")
}
