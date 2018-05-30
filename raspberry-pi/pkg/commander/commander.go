package commander

import (
	"fmt"

	"github.com/numa08/digit-hackathon/pkg/command"
)

// Commander is module of convert Command to string
type Commander interface {
	CommandToString(com command.Command) string
}

type commanderImpl struct {
}

func (c *commanderImpl) CommandToString(com command.Command) string {
	str := fmt.Sprintf("%s:%s", com.Name(), com.Value())
	return str
}

// NewCommander creates new commander instance
func NewCommander() Commander {
	return &commanderImpl{}
}
