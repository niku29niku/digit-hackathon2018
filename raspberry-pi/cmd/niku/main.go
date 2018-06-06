package main

import (
	"flag"
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/config"
	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/cooker"
	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/device"
	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/firebase"
	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/notification/phone"
	phoneRep "github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/repository/phone"
	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/repository/timer"
	pb "gopkg.in/cheggaaa/pb.v1"
	emoji "gopkg.in/kyokomi/emoji.v1"
)

func main() {
	flag.Parse()
	firebase, err := firebase.NewFirebaseClient()
	if err != nil {
		glog.Fatalf("NewFirebaseClient error : %s ", err)
	}
	configuration, err := config.DecodeDefaultToml()
	if err != nil {
		glog.Fatalf("DecodeDefaultToml error : %s ", err)
	}
	cookerDevice, err := device.GetDevice(configuration.Device)
	if err != nil {
		glog.Fatalf("GetDevice error : %s ", err)
	}
	ck := cooker.NewRoastbeefCooker(configuration.Cooker)
	err = ck.Cook(cookerDevice)
	if err != nil {
		glog.Fatalf("Cooker error: %s ", err)
	}
	timerRepository := timer.NewFirebaseRepository(firebase)
	duration := time.Duration(configuration.Cooker.Duration) * time.Second
	err = timerRepository.SetTimer(duration)
	if err != nil {
		glog.Fatalf("SetTimer error : %s ", err)
	}
	emoji.Println(":thinking: :cow: started :watch: :ok_hand:")
	bar := pb.StartNew(configuration.Cooker.Duration)
	bar.Format("🔌🍖🐃🔥📲")
	for i := 0; i < configuration.Cooker.Duration; i++ {
		bar.Increment()
		time.Sleep(1 * time.Second)
	}
	bar.FinishPrint(emoji.Sprint(":smile: :meat_on_bone: finish !! :thumbsup: :tada:"))
	err = timerRepository.Remove()
	if err != nil {
		glog.Errorf("TimerRepository.Remove error : %s ", err)
	}
	phone := phone.NewPhoneClient(configuration.Twilio, phone.NewParser())
	phoneRepository := phoneRep.NewFirebaseRepository(firebase)
	numbers, err := phoneRepository.PhoneNumbers()
	if err != nil {
		glog.Errorf("PhoneNumbers error : %s", err)
	}
	errors := phone.Notification(numbers)
	if len(errors) > 0 {
		for _, err := range errors {
			glog.Errorf("Phone Notification error : %s", err)
		}
	}
	if err != nil || len(errors) > 0 {
		os.Exit(1)
	}
	os.Exit(0)
}
