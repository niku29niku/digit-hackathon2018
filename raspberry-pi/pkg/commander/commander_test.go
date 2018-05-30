package commander

import (
	"testing"

	"github.com/numa08/digit-hackathon/pkg/command"
	"github.com/stretchr/testify/assert"
)

type MockCommand struct{}

func Test_ToString(t *testing.T) {
	com := command.Temperature(55.5)
	commander := NewCommander()
	expected := commander.CommandToString(com)
	actual := "C:55.5"
	assert.Equal(t, expected, actual)
}
