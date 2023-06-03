package adaptor

import (
	"fmt"
	"testing"
)

func TestDuckDrive(t *testing.T) {
	duck := new(MallardDuck)
	turkey := new(WildTurkey)

	turkeyAdapter := new(TurkeyAdapter)
	turkeyAdapter.Constructor(turkey)

	fmt.Println("칠면조는")
	turkey.Gobble()
	turkey.Fly()

	fmt.Println("오리는")
	duck.Quack()
	duck.Fly()

	fmt.Println("칠면조 어댑터는")
	turkeyAdapter.Quack()
	turkeyAdapter.Fly()
}
