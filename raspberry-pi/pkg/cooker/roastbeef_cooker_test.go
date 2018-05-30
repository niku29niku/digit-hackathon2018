package cooker

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/response"
)

func Test_Cook(t *testing.T) {
	t.Run("should success cooking", func(t *testing.T) {
		ctlr := gomock.NewController(t)
		defer ctlr.Finish()
		mockDevice := NewMockDevice(ctlr)
		mockDevice.EXPECT().IsReady().Return(response.Ok, nil)
		mockDevice.EXPECT().SetTemperature(gomock.Any()).Return(response.Ok, nil)
		mockDevice.EXPECT().SetDuration(gomock.Any()).Return(response.Ok, nil)
		mockDevice.EXPECT().Start().Return(response.Ok, nil)

		cooker := NewRoastbeefCooker(55.5, int(2*time.Hour))
		err := cooker.Cook(mockDevice)
		assert.Nil(t, err)
	})
	t.Run("should return error when device error", func(t *testing.T) {
		ctlr := gomock.NewController(t)
		defer ctlr.Finish()
		mockDevice := NewMockDevice(ctlr)
		mockDevice.EXPECT().IsReady().Return(response.Ng, fmt.Errorf("mock error"))

		cooker := NewRoastbeefCooker(55.5, int(2*time.Hour))
		err := cooker.Cook(mockDevice)
		expected := "mock error"
		assert.Equal(t, expected, err.Error())
	})
	t.Run("should return error when device sttatus ng", func(t *testing.T) {
		ctlr := gomock.NewController(t)
		defer ctlr.Finish()
		mockDevice := NewMockDevice(ctlr)
		mockDevice.EXPECT().IsReady().Return(response.Ng, nil)

		cooker := NewRoastbeefCooker(55.5, int(2*time.Hour))
		err := cooker.Cook(mockDevice)
		expected := "device is not ready"
		assert.Equal(t, expected, err.Error())
	})
}
