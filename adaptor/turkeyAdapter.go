package adaptor

type TurkeyAdapter struct {
	Duck
	Turkey
}

var (
	_ Turkey = (*TurkeyAdapter)(nil)
)

func (ta *TurkeyAdapter) Constructor(turkey Turkey) {
	ta.Turkey = turkey
}

func (ta *TurkeyAdapter) Quack() {
	ta.Gobble()
}

func (ta *TurkeyAdapter) Fly() {
	for i := 0; i < 5; i++ {
		ta.Turkey.Fly()
	}
}
