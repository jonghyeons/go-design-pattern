package adaptor

import "fmt"

type Turkey interface {
	Gobble()
	Fly()
}

type WildTurkey struct {
	Turkey
}

var _ Turkey = (*WildTurkey)(nil)

func (wt *WildTurkey) Gobble() {
	fmt.Println("골골")
}

func (wt *WildTurkey) Fly() {
	fmt.Println("짧은 거리를 날고 있어요")
}
