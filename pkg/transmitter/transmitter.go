// transmitter is the package to send serial message from gpio pin.
//
// The baud rate is 4800 and serial port is the primary UART of Raspberry Pi
// It looks as `/dev/ttyAMA0` from linux file system.
package transmitter

import (
	"fmt"
	"log"

	"go.bug.st/serial"
)

type Transmitter struct {
	port serial.Port
}

func (t Transmitter) Write(p []byte) (n int, err error) {
	n, err = t.port.Write(p)
	if err != nil {
		return -1, fmt.Errorf("transmitter: %v", err)
	}
	return
}

func (t Transmitter) Close() error {
	return t.port.Close()
}

func NewTransmitter() Transmitter {
	return Transmitter{initSerial()}
}

func initSerial() serial.Port {
	mode := &serial.Mode{BaudRate: 4800}
	port, err := serial.Open("/dev/ttyAMA0", mode)
	if err != nil {
		log.Fatal(err)
	}

	return port
}
