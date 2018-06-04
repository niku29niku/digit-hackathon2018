package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	t.Run("should decode config file", func(t *testing.T) {
		p, _ := filepath.Abs(filepath.Join("..", "..", "test", "config", "niku.toml"))
		_, err := os.Stat(p)
		assert.True(t, !os.IsNotExist(err))
		config, err := Decode(p)
		assert.Nil(t, err)
		twilio := config.Twilio
		assert.Equal(t, "sid", twilio.AccountSid)
		assert.Equal(t, "token", twilio.AuthToken)
		assert.Equal(t, "+810000000000", twilio.FromPhoneNumber)
		assert.Equal(t, "http://demo.twilio.com/docs/voice.xml", twilio.CallbackURL)
		cooker := config.Cooker
		assert.Equal(t, "/dev/usb.ttymodem77", cooker.DeviceName)
		assert.Equal(t, 38400, cooker.BaudRate)
	})
}
