package tv

import "github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/behavioral/command/code/interfaces"

type offCommand struct {
	device interfaces.Device
}

func NewOffCommand(device interfaces.Device) *offCommand {
	return &offCommand{
		device: device,
	}
}

func (c *offCommand) Execute() {
	c.device.Off()
}
