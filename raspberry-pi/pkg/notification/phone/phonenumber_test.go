package phone

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_E164(t *testing.T) {
	parser := NewParser()
	t.Run("should parse jp number", func(t *testing.T) {
		number := "08012345678"
		parsed, err := parser.ParseToE164(number)
		assert.Nil(t, err)
		assert.Equal(t, "+818012345678", parsed)
	})
	t.Run("should return origin number when E164 number", func(t *testing.T) {
		number := "+818012345678"
		parsed, err := parser.ParseToE164(number)
		assert.Nil(t, err)
		assert.Equal(t, "+818012345678", parsed)
	})
	t.Run("should not parse from invalid JP fone number", func(t *testing.T) {
		number := "abc"
		_, err := parser.ParseToE164(number)
		assert.Equal(t, "The phone number supplied is not a number.", err.Error())
	})
}
