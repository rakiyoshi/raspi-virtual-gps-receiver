package main

import (
	"fmt"
	"log"
	"time"

	"github.com/rakiyoshi/raspi-virtual-gps-receiver/pkg/nmealib"
	"github.com/rakiyoshi/raspi-virtual-gps-receiver/pkg/transmitter"
)

func main() {
	w := transmitter.NewTransmitter()
	defer func(w transmitter.Transmitter) {
		err := w.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(w)

	for {
		message := nmealib.FromTimeToGGA(time.Now().UTC())
		_, err := fmt.Fprint(w, message)
		if err != nil {
			log.Fatalln(err)
		}
		time.Sleep(500 * time.Millisecond)

		message = nmealib.FromTimeToRMC(time.Now().UTC())
		_, err = fmt.Fprint(w, message)
		if err != nil {
			log.Fatalln(err)
		}
		time.Sleep(500 * time.Millisecond)
	}
}
