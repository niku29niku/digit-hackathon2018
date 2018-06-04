package main

import (
	"fmt"
	"os"

	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/config"
	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/cooker"
	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/device"
	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/notification/phone"
)

func main() {
	configuration, err := config.DecodeDefaultToml()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Configuration error: %s \n", err.Error())
		os.Exit(1)
	}
	cookerDevice, err := device.GetDevice(configuration.Device)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cookerdevice error: %s \n", err.Error())
		os.Exit(1)
	}
	ck := cooker.NewRoastbeefCooker(configuration.Cooker)
	err = ck.Cook(cookerDevice)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cooker error: %s \n", err.Error())
		os.Exit(1)
	}
	phone := phone.NewPhoneClient(configuration.Twilio, phone.NewParser())
	errors := phone.Notification([]string{"080-5238-6255"})
	if len(errors) > 0 {
		for _, err := range errors {
			fmt.Fprintf(os.Stderr, "Phone error: %s \n", err.Error())
		}
		os.Exit(1)
	}
	os.Exit(0)
}
