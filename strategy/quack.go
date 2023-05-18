package strategy

import "fmt"

type QuackBehavior interface {
	Quack()
}

type Quack struct {
	QuackBehavior
}

func (q *Quack) Quack() {
	fmt.Println("꽥꽥")
}

type Squeak struct {
	QuackBehavior
}

func (s *Squeak) Quack() {
	fmt.Println("삑삑")
}

type MuteQuack struct {
	QuackBehavior
}

func (mq *MuteQuack) Quack() {
	fmt.Println()
}
