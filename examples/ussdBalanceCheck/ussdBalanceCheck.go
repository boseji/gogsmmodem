// Example of Check Balance using USSD Code.
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

	val, err := modem.SendUSSDSingle("*123#")
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n\nReceived \n%#v\n\n", val)
}
