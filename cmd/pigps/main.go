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
		fmt.Print(message)
		time.Sleep(1 * time.Second)
	}
}
