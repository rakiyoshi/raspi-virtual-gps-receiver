package main

import (
	"fmt"
	"time"

	"github.com/rakiyoshi/raspi-virtual-gps-receiver/pkg/nmealib"
	"github.com/rakiyoshi/raspi-virtual-gps-receiver/pkg/transmitter"
)

func main() {
	w := transmitter.NewTransmitter()
	defer w.Close()

	for {
		message := nmealib.FromTimeToGGA(time.Now().UTC())
		fmt.Fprint(w, message)
		time.Sleep(500 * time.Millisecond)

		message = nmealib.FromTimeToRMC(time.Now().UTC())
		fmt.Fprint(w, message)
		time.Sleep(500 * time.Millisecond)
	}
}
