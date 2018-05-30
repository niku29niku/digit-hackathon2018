package cooker

import (
	"fmt"

	"github.com/numa08/digit-hackathon2018/raspberry-pi/pkg/device"
	"github.com/numa08/digit-hackathon2018/raspberry-pi/pkg/response"
)

type roastbeefCooker struct {
	temperture float64
	duration   int
}

func (ck *roastbeefCooker) Cook(device device.Device) (err error) {
	var status response.Status
	status, err = device.IsReady()
	if err != nil {
		return
	}
	if status != response.Ok {
		err = fmt.Errorf("device is not ready")
		return
	}
	status, err = device.SetTemperature(ck.temperture)
	if err != nil {
		return
	}
	if status != response.Ok {
		err = fmt.Errorf("failed set temperature")
		return
	}
	status, err = device.SetDuration(ck.duration)
	if err != nil {
		return
	}
	if status != response.Ok {
		err = fmt.Errorf("failed set duration")
		return
	}
	status, err = device.Start()
	if err != nil {
		return
	}
	if status != response.Ok {
		err = fmt.Errorf("failed start")
		return
	}
	return nil
}
