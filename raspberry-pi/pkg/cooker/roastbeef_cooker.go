package cooker

import (
	"fmt"

	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/device"
	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/response"
)

type roastbeefCooker struct {
	temperture float64
	duration   int
}

func (ck *roastbeefCooker) Cook(device device.Device) (err error) {
	serial := func(f func() (response.Status, error), msg string) (st response.Status) {
		if err != nil {
			return
		}
		st, err = f()
		if err != nil {
			return
		}
		if st != response.Ok {
			err = fmt.Errorf(msg)
			return
		}
		return st
	}
	_ = serial(func() (response.Status, error) { return device.IsReady() }, "device is not ready")
	_ = serial(func() (response.Status, error) { return device.SetTemperature(ck.temperture) }, "failed set temperature")
	_ = serial(func() (response.Status, error) { return device.SetDuration(ck.duration) }, "failed set duration")
	_ = serial(func() (response.Status, error) { return device.Start() }, "failed start")
	return
}
