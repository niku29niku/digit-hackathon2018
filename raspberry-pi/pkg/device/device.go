package device

import (
	"regexp"
	"sync"
	"time"

	"github.com/golang/glog"
	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/command"
	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/commander"
	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/config"
	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/response"
	"github.com/tarm/serial"
)

// Device is serial connection device
type Device interface {
	SetTemperature(value float64) (response.Status, error)
	SetDuration(value int) (response.Status, error)
	IsReady() (response.Status, error)
	Start() (response.Status, error)
}

type arduino struct {
	port      *serial.Port
	commander commander.Commander
	parser    response.Parser
}

var sharedArduinoInstance *arduino
var once sync.Once

const timeoutDuration = 5 * time.Second

// GetDevice get Device instance
func GetDevice(config config.DeviceConfig) (dev Device, err error) {
	once.Do(func() {
		config := &serial.Config{Name: config.DeviceName, Baud: config.BaudRate, ReadTimeout: timeoutDuration}
		glog.V(2).Infof("Address : %s, Baudrate : %d ", config.Name, config.Baud)
		port, e := serial.OpenPort(config)
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
	wr := []byte(com)
	glog.V(2).Infof("write command %s", wr)
	_, err = dev.port.Write([]byte(com))
	if err != nil {
		return
	}
	b := make([]byte, 10, 10)
	_, err = dev.port.Read(b)
	if err != nil {
		return
	}
	glog.V(2).Infof("read result %s", b)
	s := string(b)
	re := regexp.MustCompile("(.+)\n")
	resp = re.FindAllStringSubmatch(s, -1)[0][0]
	return
}
