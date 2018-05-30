package device

import (
	"math"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/numa08/digit-hackathon2018/raspberry-pi/pkg/commander"
	"github.com/numa08/digit-hackathon2018/raspberry-pi/pkg/response"
	"github.com/stretchr/testify/assert"
)

func Test_StatusCommand(t *testing.T) {
	ctlr := gomock.NewController(t)
	defer ctlr.Finish()
	mockSerial := NewMockPort(ctlr)
	mockSerial.EXPECT().Write(gomock.Any()).Return(0, nil).MaxTimes(math.MaxInt64)
	t.Run("should get Ok when return ok status", func(t *testing.T) {
		mockSerial.EXPECT().Read(gomock.Any()).DoAndReturn(func(p []byte) (int, error) {
			p[0] = 'O'
			p[1] = 'K'
			p[2] = '\r'
			p[3] = '\n'
			return len(p), nil
		})

		device := &arduino{mockSerial, commander.NewCommander(), response.NewParser()}
		status, err := device.SetTemperature(55.5)
		assert.Equal(t, status, response.Ok)
		assert.Nil(t, err)
	})
	t.Run("should get Ng when return ng status", func(t *testing.T) {
		mockSerial.EXPECT().Read(gomock.Any()).DoAndReturn(func(p []byte) (int, error) {
			p[0] = 'N'
			p[1] = 'G'
			p[2] = '\r'
			p[3] = '\n'
			return len(p), nil
		})
		device := &arduino{mockSerial, commander.NewCommander(), response.NewParser()}
		status, err := device.SetDuration(600)
		assert.Equal(t, status, response.Ng)
		assert.Nil(t, err)
	})
}

func Test_FaildConnection(t *testing.T) {
	t.Run("should get error when return unformatted status", func(t *testing.T) {
		ctlr := gomock.NewController(t)
		defer ctlr.Finish()
		mockSerial := NewMockPort(ctlr)
		mockSerial.EXPECT().Write(gomock.Any()).Return(0, nil).MaxTimes(math.MaxInt64)
		mockSerial.EXPECT().Read(gomock.Any()).DoAndReturn(func(p []byte) (int, error) {
			p[0] = 'K'
			p[1] = 'O'
			p[2] = '\r'
			p[3] = '\n'
			return len(p), nil
		})
		device := &arduino{mockSerial, commander.NewCommander(), response.NewParser()}
		_, err := device.IsReady()
		assert.Equal(t, err.Error(), "response is invalid KO\r\n")
	})
	t.Run("should get error when write timeout", func(t *testing.T) {
		ctlr := gomock.NewController(t)
		defer ctlr.Finish()
		mockSerial := NewMockPort(ctlr)
		mockSerial.EXPECT().Write(gomock.Any()).DoAndReturn(func(p []byte) (int, error) {
			time.Sleep(10 * time.Second)
			return len(p), nil
		})
		device := &arduino{mockSerial, commander.NewCommander(), response.NewParser()}
		_, err := device.Start()
		assert.Equal(t, err.Error(), "write timeout for 5 second")
	})
	t.Run("shoud get error when read timeout", func(t *testing.T) {
		ctlr := gomock.NewController(t)
		defer ctlr.Finish()
		mockSerial := NewMockPort(ctlr)
		mockSerial.EXPECT().Write(gomock.Any()).Return(0, nil).MaxTimes(math.MaxInt64)
		mockSerial.EXPECT().Read(gomock.Any()).DoAndReturn(func(p []byte) (int, error) {
			time.Sleep(10 * time.Second)
			return 0, nil
		})
		device := &arduino{mockSerial, commander.NewCommander(), response.NewParser()}
		_, err := device.Start()
		assert.Equal(t, err.Error(), "read timeout for 5 second")
	})
}
