package config

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/golang/glog"
)

// Config is defined configuration file
type Config struct {
	Twilio TwilioConfig `toml:"twilio"`
	Cooker CookerConfig `toml:"cooker"`
	Device DeviceConfig `toml:"device"`
}

// TwilioConfig is defined to Twilio configuration
type TwilioConfig struct {
	AccountSid      string `toml:"sid"`
	AuthToken       string `toml:"token"`
	FromPhoneNumber string `toml:"phone_number"`
	CallbackURL     string `toml:"callback_url"`
}

// DeviceConfig is defined to device configuration
type DeviceConfig struct {
	DeviceName string `toml:"device"`
	BaudRate   int    `toml:"baudrate"`
}

// CookerConfig is defined to cooking configuration
type CookerConfig struct {
	Duration   int     `toml:"duration"`
	Temperture float64 `toml:"temperture"`
}

// Decode config file to Config
func Decode(path string) (config Config, err error) {
	_, err = toml.DecodeFile(path, &config)
	return config, err
}

// DecodeDefaultToml will decode config file in $HOME/niku.toml
func DecodeDefaultToml() (Config, error) {
	path := filepath.Join(os.Getenv("HOME"), "niku.toml")
	glog.V(2).Infof("configuration file path : %s", path)
	return Decode(path)
}
