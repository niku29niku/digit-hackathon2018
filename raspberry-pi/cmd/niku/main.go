package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/config"
	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/cooker"
	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/device"
	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/firebase"
	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/notification/phone"
	phoneRep "github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/repository/phone"
	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/repository/timer"
)

func main() {
	firebase, err := firebase.NewFirebaseClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "NewFirebaseClient error : %s \n", err.Error())
		os.Exit(1)
	}
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
	timerRepository := timer.NewFirebaseRepository(firebase)
	err = timerRepository.SetTimer(time.Duration(configuration.Cooker.Duration) * time.Second)
	if err != nil {
		fmt.Fprintf(os.Stderr, "TimerRepository error : %s \n", err.Error())
		os.Exit(1)
	}
	err = timerRepository.Remove()
	if err != nil {
		fmt.Fprintf(os.Stderr, "TimerRepository error : %s \n", err.Error())
		os.Exit(1)
	}

	phone := phone.NewPhoneClient(configuration.Twilio, phone.NewParser())
	phoneRepository := phoneRep.NewFirebaseRepository(firebase)
	numbers, err := phoneRepository.PhoneNumbers()
	if err != nil {
		log.Fatalf("PhoneNumbers error : %s", err.Error())
	}
	errors := phone.Notification(numbers)
	if len(errors) > 0 {
		for _, err := range errors {
			fmt.Fprintf(os.Stderr, "Phone error: %s \n", err.Error())
		}
		os.Exit(1)
	}
	os.Exit(0)
}
