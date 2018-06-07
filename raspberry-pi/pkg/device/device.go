package device

import (
	"bytes"
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

const timeoutDuration = 10 * time.Second

// GetDevice get Device instance
func GetDevice(config config.DeviceConfig) (dev Device, err error) {
	cfg := &serial.Config{Name: config.DeviceName, Baud: config.BaudRate, ReadTimeout: timeoutDuration}
	glog.V(2).Infof("Address : %s, Baudrate : %d ", cfg.Name, cfg.Baud)
	port, e := serial.OpenPort(cfg)
	if e != nil {
		err = e
		return
	}
	sharedArduinoInstance = &arduino{
		commander: commander.NewCommander(),
		port:      port,
		parser:    response.NewParser(),
	}
	dev = sharedArduinoInstance
	time.Sleep(5 * time.Second)
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
	b := make([]byte, 0)
	// 1回の Read でデータの読み込みができないことがあるのでループする
	for {
		if bytes.Contains(b, []byte("\r\n")) {
			break
		}
		nb := make([]byte, 10, 10)
		_, err = dev.port.Read(nb)
		if err != nil {
			return
		}
		b = append(b, nb...)
	}
	// 通信ノイズ？で 0 が含まれることがあるので消去する
	b = filter(b, func(e byte) bool { return e != 0 })
	for _, i := range b {
		glog.V(2).Infof("read byte %d", i)
	}
	glog.V(2).Infof("read result %s", b)
	resp = string(b)
	for _, i := range []byte(resp) {
		glog.V(2).Infof("reasult string byte %d", i)
	}
	return
}

func filter(vs []byte, f func(byte) bool) []byte {
	vsf := make([]byte, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
