package command

import "testing"

func TestCommand(t *testing.T) {
	controller := Controller{}

	controller.SetCommand(LightOnCommand{})
	controller.Call()

	controller.SetCommand(LightOffCommand{})
	controller.Call()
}
