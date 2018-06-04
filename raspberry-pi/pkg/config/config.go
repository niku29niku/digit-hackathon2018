package config

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// Config is defined configuration file
type Config struct {
	Twilio TwilioConfig `toml:"twilio"`
	Cooker CookerConfig `toml:"cooker"`
}

// TwilioConfig is defined to Twilio configuration
type TwilioConfig struct {
	AccountSid      string `toml:"sid"`
	AuthToken       string `toml:"token"`
	FromPhoneNumber string `toml:"phone_number"`
	CallbackURL     string `toml:"callback_url"`
}

// CookerConfig is defined to cooker device configuration
type CookerConfig struct {
	DeviceName string `toml:"device"`
	BaudRate   int    `toml:"baudrate"`
}

// Decode config file to Config
func Decode(path string) (config Config, err error) {
	_, err = toml.DecodeFile(path, &config)
	return config, err
}

// DecodeDefaultToml will decode config file in $HOME/niku.toml
func DecodeDefaultToml() (Config, error) {
	path := filepath.Join(os.Getenv("HOME"), "niku.toml")
	return Decode(path)
}
