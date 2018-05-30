package response

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParseStatus(t *testing.T) {
	parser := NewParser()
	t.Run("should get ok when response string is `ok`", func(t *testing.T) {
		st, err := parser.ParseStatus("ok\r\n")
		assert.Nil(t, err)
		assert.Equal(t, st, Ok)
	})
	t.Run("should get ng when response string is `ng`", func(t *testing.T) {
		st, err := parser.ParseStatus("ng\r\n")
		assert.Nil(t, err)
		assert.Equal(t, st, Ng)
	})
	t.Run("should get error when response string is not supported", func(t *testing.T) {
		_, err := parser.ParseStatus("undefined\r\n")
		assert.Equal(t, err.Error(), "response is invalid undefined\r\n")
	})
}
