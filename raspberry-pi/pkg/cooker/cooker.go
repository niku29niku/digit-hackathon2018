package cooker

import "github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/device"

// Cooker cook
type Cooker interface {
	Cook(device device.Device) error
}

// NewRoastbeefCooker create new cooker instance for roastbeef
func NewRoastbeefCooker(temperture float64, duration int) Cooker {
	return &roastbeefCooker{
		duration:   duration,
		temperture: temperture,
	}
}
