// Example of Sending a SMS.
package main

import (
	"fmt"

	"github.com/boseji/gogsmmodem"
	"github.com/tarm/serial"
)

func main() {
	conf := serial.Config{Name: "/dev/ttyUSB0", Baud: 9600}
	modem, err := gogsmmodem.Open(&conf, true)
	if err != nil {
		panic(err)
	}

	err = modem.SendMessage("+919845098450", "Hari Aum")
	if err != nil {
		panic(err)
	}
	fmt.Print("SMS Sent\n\n")
}
