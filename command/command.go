package command

type Command interface {
	Execute()
}

type LightOnCommand struct {
	light *Light
	Command
}

var _ Command = (*LightOnCommand)(nil)

func (loc *LightOnCommand) Constructor(light *Light) {
	loc.light = light

}
func (loc *LightOnCommand) Execute() {
	loc.light.On()
}
