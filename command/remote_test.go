package command

import "testing"

func TestRemoteControl(t *testing.T) {
	remote := new(SimpleRemoteControl)
	light := new(Light)
	lightOn := LightOnCommand{}
	lightOn.Constructor(light)

	remote.SetCommand(&lightOn)
	remote.ButtonWasPressed()
}
