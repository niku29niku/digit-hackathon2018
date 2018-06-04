package cooker

import (
	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/config"
	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/device"
)

// Cooker cook
type Cooker interface {
	Cook(device device.Device) error
}

// NewRoastbeefCooker create new cooker instance for roastbeef
func NewRoastbeefCooker(config config.CookerConfig) Cooker {
	return &roastbeefCooker{
		duration:   config.Duration,
		temperture: config.Temperture,
	}
}
