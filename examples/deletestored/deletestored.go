// Example of retrieving and deleting stored messages.
package main

import (
	"fmt"

	"github.com/barnybug/gogsmmodem"
	"github.com/tarm/serial"
)

func main() {
	conf := serial.Config{Name: "/dev/ttyUSB0", Baud: 9600}
	modem, err := gogsmmodem.Open(&conf, true)
	if err != nil {
		panic(err)
	}

	li, _ := modem.ListMessages("ALL")
	fmt.Printf("%d stored messages\n", len(*li))
	for _, msg := range *li {
		fmt.Println(msg.Index, msg.Status, msg.Body)
		err := modem.DeleteMessage(msg.Index)
		if err != nil {
			panic(err)
		}
		fmt.Println("Deleted Message ", msg.Index)
	}
}
