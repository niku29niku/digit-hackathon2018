package phone

import (
	"fmt"
	"testing"

	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/notification/phone/twilio"
	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

func Test_Notification(t *testing.T) {
	t.Run("should get no error valid phone numbers", func(t *testing.T) {
		ctlr := gomock.NewController(t)
		defer ctlr.Finish()
		mockTwilio := twilio.NewMockTwilio(ctlr)
		mockTwilio.EXPECT().Call(gomock.Any()).Return(nil).Times(3)
		numbers := []string{
			"080-4972-271",
			"080-9962-9297",
			"080-9852-2523",
		}

		client := twilioClient{
			Twilio:       mockTwilio,
			NumberParser: NewParser(),
		}
		errors := client.Notification(numbers)
		assert.Empty(t, errors)
	})
	t.Run("should get one error when invalid phone number", func(t *testing.T) {
		ctlr := gomock.NewController(t)
		defer ctlr.Finish()
		mockTwilio := twilio.NewMockTwilio(ctlr)
		mockTwilio.EXPECT().Call(gomock.Any()).Return(nil).Times(2)
		numbers := []string{
			"abcd",
			"080-9962-9297",
			"080-9852-2523",
		}

		client := twilioClient{
			Twilio:       mockTwilio,
			NumberParser: NewParser(),
		}
		errors := client.Notification(numbers)
		assert.Equal(t, 1, len(errors))
		assert.Equal(t, "abcd, error: The phone number supplied is not a number.", errors[0].Error())
	})
	t.Run("should get one error when twilio call failed", func(t *testing.T) {
		ctlr := gomock.NewController(t)
		defer ctlr.Finish()
		mockTwilio := twilio.NewMockTwilio(ctlr)
		mockTwilio.EXPECT().Call(gomock.Any()).DoAndReturn(func(num string) error {
			if num == "+81804972271" {
				return fmt.Errorf("twilio failed")
			}
			return nil
		}).Times(3)
		numbers := []string{
			"080-4972-271",
			"080-9962-9297",
			"080-9852-2523",
		}

		client := twilioClient{
			Twilio:       mockTwilio,
			NumberParser: NewParser(),
		}
		errors := client.Notification(numbers)
		assert.Equal(t, 1, len(errors))
		assert.Equal(t, "080-4972-271, error: twilio failed", errors[0].Error())
	})
}
