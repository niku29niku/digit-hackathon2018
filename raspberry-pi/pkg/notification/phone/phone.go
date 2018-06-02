package phone

import (
	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/config"
	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/notification/phone/twilio"
)

// Phone is interface to phone call notification
type Phone interface {
	Notification(to []string) []error
}

// NewPhoneClient create new instance
func NewPhoneClient(
	config config.TwilioConfig, numberParser NumberParser) Phone {
	client := twilio.NewTwilioClient(config)
	return &twilioClient{
		Twilio:       client,
		NumberParser: numberParser,
	}
}
