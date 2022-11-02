package main

import (
	"fmt"
	"time"

	"github.com/rakiyoshi/raspi-virtual-gps-receiver/pkg/transmitter"
)

func main() {
	w := transmitter.NewTransmitter()
	defer w.Close()

	for {
		// Deprecated: this is dummy
		dummy := "$GPZDA,092403,15,07,2017,,"

		fmt.Fprintln(w, dummy)
		time.Sleep(10 * time.Second)
	}
}
