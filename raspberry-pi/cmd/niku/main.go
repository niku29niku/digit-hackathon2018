package main

import (
	"fmt"
	"os"
	"time"

	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/cooker"
	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/device"
)

func main() {
	d, err := device.GetDevice("/dev/tty.usbmodem77", 38400)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed get device : %s\n", err.Error())
		os.Exit(1)
	}
	duration := int(2 * time.Hour / time.Second)
	ck := cooker.NewRoastbeefCooker(55.5, duration)
	err = ck.Cook(d)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err.Error())
		os.Exit(1)
	}
}
