# go-cantact

A package for communicating on Controller Area Network buses using the
[CANtact](http://cantact.io) hardawre.

## Details

This package makes use of the [serial](https://github.com/tarm/serial) library
for providing access to CANtact's serial port. It provides functions for sending
and receiving frames in the correct format for the device.

## Usage
```go
package main

import (
	"log"
	"github.com/linklayer/go-cantact"
)

func main() {
	d, err := cantact.NewDevice(os.Args[1])
	if err != nil:
		log.Fatal(err)

	// set bitrate mode to 6, 500 kbps
	d.SetBitrate(6)

	// open connection to CAN bus
	d.Open()

	// send a frame
	data := []byte{0xDE, 0xAD, 0xBE, 0xEF}
	tx_frame := Frame{Id: 0x123, Dlc: len(data), Data: data}
	d.WriteFrame(tx_frame)

	// read a frame (blocking call)
	rx_frame := d.ReadFrame()
	log.Println(rx_frame)

}
```

