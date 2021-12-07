package tv

import "github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/command/code/interfaces"

type onCommand struct {
	device interfaces.Device
}

func NewOnCommand(device interfaces.Device) *onCommand {
	return &onCommand{
		device: device,
	}
}

func (c *onCommand) Execute() {
	c.device.On()
}