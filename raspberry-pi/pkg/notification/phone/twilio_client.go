package phone

import (
	"fmt"

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
			e = t.Twilio.Call(num)
			if e != nil {
				channel <- fmt.Errorf("%s, error: %s", toNum, e.Error())
				return
			}
			channel <- nil
			return
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
