package strategy

import "testing"

func TestDuck(t *testing.T) {
	// 물오리
	mallardDuck := &Duck{
		FlyBehavior:   &FlyWithWings{},
		QuackBehavior: &Quack{},
	}

	mallardDuck.Fly()
	mallardDuck.Quack()
}
