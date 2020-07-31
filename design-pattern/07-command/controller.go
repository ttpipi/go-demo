package command

type Controller struct {
	cm Command
}

func (c *Controller) SetCommand(cm Command) {
	c.cm = cm
}

func (c *Controller) Call() {
	c.cm.Execute()
}
