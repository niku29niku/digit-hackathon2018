package commander

import (
	"testing"

	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/command"
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
