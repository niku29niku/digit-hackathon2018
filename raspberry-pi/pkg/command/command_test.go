package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TempratureCommand(t *testing.T) {
	cmd := Temperature(55.5)
	t.Run("name", func(t *testing.T) {
		actual := cmd.Name()
		expected := "C"
		assert.Equal(t, actual, expected)
	})
	t.Run("value", func(t *testing.T) {
		actual := cmd.Value()
		expected := "55.5"
		assert.Equal(t, actual, expected)
	})
}

func Test_Duration(t *testing.T) {
	cmd := Duration(600)
	t.Run("name", func(t *testing.T) {
		actual := cmd.Name()
		expected := "T"
		assert.Equal(t, actual, expected)
	})
	t.Run("value", func(t *testing.T) {
		actual := cmd.Value()
		expected := "600"
		assert.Equal(t, actual, expected)
	})
}

func Test_Start(t *testing.T) {
	cmd := Start()
	t.Run("name", func(t *testing.T) {
		actual := cmd.Name()
		expected := "L"
		assert.Equal(t, actual, expected)
	})
	t.Run("value", func(t *testing.T) {
		actual := cmd.Value()
		expected := ""
		assert.Equal(t, actual, expected)
	})
}

func Test_Stop(t *testing.T) {
	cmd := Stop()
	t.Run("name", func(t *testing.T) {
		actual := cmd.Name()
		expected := "E"
		assert.Equal(t, actual, expected)
	})
	t.Run("value", func(t *testing.T) {
		actual := cmd.Value()
		expected := ""
		assert.Equal(t, actual, expected)
	})
}

func Test_Ready(t *testing.T) {
	cmd := Ready()
	t.Run("name", func(t *testing.T) {
		actual := cmd.Name()
		expected := "?"
		assert.Equal(t, actual, expected)
	})
	t.Run("value", func(t *testing.T) {
		actual := cmd.Value()
		expected := "L"
		assert.Equal(t, actual, expected)
	})
}
