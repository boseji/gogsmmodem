## gogsmmodem: Go library for GSM modems

[![Build Status](https://travis-ci.org/barnybug/gogsmmodem.svg?branch=master)](https://travis-ci.org/barnybug/gogsmmodem)

Go library for the sending and receiving SMS messages through a GSM modem.

#### Update January 15th, 2019 

Now, also support USSD requests for balance enquiry and call forwarding. 

And phone status Check.

### Tested devices
- ZTE MF110/MF627/MF636
- SIM900A, SIM900, SIM300z (Module over Serial)

### Installation
Run:

    go get github.com/barnybug/gogsmmodem

### Usage
Example:

```go

package main

import (
    "fmt"

    "github.com/barnybug/gogsmmodem"
    "github.com/tarm/serial"
)

func main() {
    conf := serial.Config{Name: "/dev/ttyUSB1", Baud: 115200}
    modem, err := gogsmmodem.Open(&conf, true)
    if err != nil {
        panic(err)
    }

    for packet := range modem.OOB {
        fmt.Printf("%#v\n", packet)
        switch p := packet.(type) {
        case gogsmmodem.MessageNotification:
            fmt.Println("Message notification:", p)
            msg, err := modem.GetMessage(p.Index)
            if err == nil {
                fmt.Printf("Message from %s: %s\n", msg.Telephone, msg.Body)
                modem.DeleteMessage(p.Index)
            }
        }
    }
}
```

### `Module` Usage Guide for this library

If, you are using a module such as SIM900 or SIM300z, <br>
then the **BEFORE EXECUTING your PROGRAM** make sure of :

- Module is connected to the PC / Board running the `golang` program
- Module is Powered properly
- SIM card is present in the Module
- Module has connected over to the network

All the the above steps need to be true for this package to work properly.

### Changelog
0.1.0

- First release

