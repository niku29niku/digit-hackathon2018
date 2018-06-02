package twilio

import (
	"fmt"

	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/config"

	"github.com/sfreiberg/gotwilio"
)

// Twilio interface is gotwilio wrapper
type Twilio interface {
	Call(number string) error
}

// NewTwilioClient create new Twilio instance
func NewTwilioClient(config config.TwilioConfig) Twilio {
	cli := gotwilio.NewTwilioClient(config.AccountSid, config.AuthToken)
	callback := gotwilio.NewCallbackParameters(config.CallbackURL)
	fromNumber := config.FromPhoneNumber
	return &twilioImpl{
		callbackParams: callback,
		twilioClient:   cli,
		fromNumber:     fromNumber,
	}
}

type twilioImpl struct {
	callbackParams *gotwilio.CallbackParameters
	twilioClient   *gotwilio.Twilio
	fromNumber     string
}

func (t *twilioImpl) Call(to string) error {
	_, exp, err := t.twilioClient.CallWithUrlCallbacks(t.fromNumber, to, t.callbackParams)
	if exp != nil || err != nil {
		return fmt.Errorf("exception: %s, error: %s", exp.Message, err.Error())
	}
	return nil
}
