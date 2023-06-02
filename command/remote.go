package command

type SimpleRemoteControl struct {
	slot *Command
}

func (src *SimpleRemoteControl) SetCommand(command Command) {
	src.slot = &command
}

func (src *SimpleRemoteControl) ButtonWasPressed() {
	(*src.slot).Execute()
}
