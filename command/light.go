package command

import "fmt"

type Light struct{}

func (l *Light) On() {
	fmt.Println("조명을 킵니다")
}
func (l *Light) Off() {
	fmt.Println("조명을 끕니다")
}
