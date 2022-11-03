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
		fmt.Fprintln(w, nmealib.FromTime(time.Now().UTC()))
		time.Sleep(1 * time.Second)
	}
}
