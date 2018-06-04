package phone

import (
	"fmt"
	"time"

	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/notification/phone/twilio"
)

type twilioClient struct {
	Twilio       twilio.Twilio
	NumberParser NumberParser
}

func (t *twilioClient) Notification(to []string) []error {
	errors := t.call(to)
	return errors
}

func (t *twilioClient) call(tos []string) []error {
	channel := make(chan error)
	for _, to := range tos {
		go func(toNum string) {
			num, e := t.NumberParser.ParseToE164(toNum)
			if e != nil {
				channel <- fmt.Errorf("%s, error: %s", toNum, e.Error())
				return
			}
			c := make(chan error)
			go func(n string, c chan error) {
				c <- t.Twilio.Call(n)
			}(num, c)
			for {
				select {
				case receive := <-c:
					if receive != nil {
						e = fmt.Errorf("%s, error: %s", toNum, receive.Error())
					}
					channel <- e
					return
				case <-time.After(5 * time.Second):
					e = fmt.Errorf("%s, error: twilio timeout", toNum)
					channel <- e
					return
				}
			}
		}(to)
	}
	errors := make([]error, 0)
	for range tos {
		err := <-channel
		if err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}
