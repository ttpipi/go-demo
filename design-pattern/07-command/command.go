package command

type Command interface {
	Execute()
}

type LightOnCommand struct {
	Light
}

func (l LightOnCommand) Execute() {
	l.On()
}

type LightOffCommand struct {
	Light
}

func (l LightOffCommand) Execute() {
	l.Off()
}
