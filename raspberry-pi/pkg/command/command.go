package command

import (
	"fmt"
)

// Command means command to serial prot
type Command interface {
	Name() string
	Value() string
}

type temperatureCommand struct {
	value float64
}

func (c *temperatureCommand) Name() string {
	return "C"
}

func (c *temperatureCommand) Value() string {
	return fmt.Sprintf("%.1f", c.value)
}

type durationCommand struct {
	value int
}

func (c *durationCommand) Name() string {
	return "T"
}

func (c *durationCommand) Value() string {
	return fmt.Sprintf("%d", c.value)
}

type startCommand struct{}

func (c *startCommand) Name() string {
	return "L"
}

func (c *startCommand) Value() string {
	return ""
}

type stopCommand struct{}

func (c *stopCommand) Name() string {
	return "E"
}

func (c *stopCommand) Value() string {
	return ""
}

// Temperature create temperature command
func Temperature(value float64) Command {
	return &temperatureCommand{
		value: value,
	}
}

// Duration create duration command
func Duration(value int) Command {
	return &durationCommand{
		value: value,
	}
}

// Start create start command
func Start() Command {
	return &startCommand{}
}

// Stop create stop commdnd
func Stop() Command {
	return &stopCommand{}
}

type readyCommand struct{}

func (c *readyCommand) Name() string {
	return "?"
}

func (c *readyCommand) Value() string {
	return "L"
}

// Ready create cook ready command
func Ready() Command {
	return &readyCommand{}
}
