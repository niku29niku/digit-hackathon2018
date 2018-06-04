package device

import (
	"fmt"
	"regexp"
	"sync"
	"time"

	"github.com/goburrow/serial"
	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/command"
	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/commander"
	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/config"
	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/response"
)

// Device is serial connection device
type Device interface {
	SetTemperature(value float64) (response.Status, error)
	SetDuration(value int) (response.Status, error)
	IsReady() (response.Status, error)
	Start() (response.Status, error)
}

type arduino struct {
	port      serial.Port
	commander commander.Commander
	parser    response.Parser
}

var sharedArduinoInstance *arduino
var once sync.Once

const timeoutDuration = 5 * time.Second

// GetDevice get Device instance
func GetDevice(config config.DeviceConfig) (dev Device, err error) {
	once.Do(func() {
		config := &serial.Config{Address: config.DeviceName, BaudRate: config.BaudRate}
		port, e := serial.Open(config)
		if e != nil {
			err = e
			return
		}
		sharedArduinoInstance = &arduino{
			commander: commander.NewCommander(),
			port:      port,
			parser:    response.NewParser(),
		}
	})
	dev = sharedArduinoInstance
	return
}

func (dev *arduino) SetTemperature(value float64) (response.Status, error) {
	com := command.Temperature(value)
	return dev.executeStatusCommand(com)
}

func (dev *arduino) SetDuration(value int) (response.Status, error) {
	com := command.Duration(value)
	return dev.executeStatusCommand(com)
}

func (dev *arduino) IsReady() (response.Status, error) {
	com := command.Ready()
	return dev.executeStatusCommand(com)
}

func (dev *arduino) Start() (response.Status, error) {
	com := command.Start()
	return dev.executeStatusCommand(com)
}

func (dev *arduino) executeStatusCommand(com command.Command) (st response.Status, er error) {
	str := dev.commander.CommandToString(com)
	res, er := dev.executeCommand(str)
	if er != nil {
		return
	}
	st, er = dev.parser.ParseStatus(res)
	return
}

func (dev *arduino) executeCommand(com string) (resp string, err error) {
	result := make(chan readResult)
	go func() {
		_, e := dev.port.Write([]byte(com))
		if e != nil {
			result <- readResult{"", e}
			return
		}
		b := make([]byte, 10, 10)
		_, e = dev.port.Read(b)
		s := string(b)
		re := regexp.MustCompile("(.+)\n")
		rs := re.FindAllStringSubmatch(s, -1)
		result <- readResult{rs[0][0], e}
	}()
	for {
		select {
		case receive := <-result:
			resp = receive.str
			err = receive.err
			return
		case <-time.After(timeoutDuration):
			err = fmt.Errorf("connection timeout for %d second", timeoutDuration/time.Second)
			return
		}
	}
}

type readResult struct {
	str string
	err error
}
