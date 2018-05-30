package device

import (
	"bufio"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/goburrow/serial"
	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/command"
	"github.com/niku29niku/digit-hackathon2018/raspberry-pi/pkg/commander"
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
func GetDevice(address string, baudRate int) (dev Device, err error) {
	once.Do(func() {
		config := &serial.Config{Address: address, BaudRate: baudRate}
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
	_, err = asyncWrite(dev.port, []byte(com))
	if err != nil {
		return
	}
	resp, err = asyncRead(bufio.NewReader(dev.port), '\n')
	return
}

func asyncWrite(writer io.Writer, b []byte) (nn int, err error) {
	wr := make(chan writeResult)
	go func() {
		n, e := writer.Write(b)
		wr <- writeResult{n, e}
	}()
	for {
		select {
		case receive := <-wr:
			nn = receive.nn
			err = receive.err
			return
		case <-time.After(timeoutDuration):
			err = fmt.Errorf("write timeout for %d second", timeoutDuration/time.Second)
			return
		}
	}
}

func asyncRead(reader *bufio.Reader, delim byte) (str string, err error) {
	rr := make(chan readResult)
	go func() {
		s, e := reader.ReadString(delim)
		rr <- readResult{s, e}
	}()
	for {
		select {
		case receive := <-rr:
			str = receive.str
			err = receive.err
			return
		case <-time.After(timeoutDuration):
			err = fmt.Errorf("read timeout for %d second", timeoutDuration/time.Second)
			return
		}
	}
}

type writeResult struct {
	nn  int
	err error
}

type readResult struct {
	str string
	err error
}
