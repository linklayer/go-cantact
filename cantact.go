// Copyright 2016 Eric Evenchick. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file

/*
Package cantact implements an interface to the CANtact device, to provide 
to Controller Area Network (CAN) buses.
*/
package cantact

import (
	"github.com/tarm/serial"
	"fmt"
)

// Frame is a single CAN frame.
type Frame struct {
	ID int
	Dlc int
	Data []byte
}

// Device is a reference to a hardware CANtact device connected to this
// computer.
type Device struct {
	port *serial.Port
}

// NewDevice creates a new device.
func NewDevice(portName string) (Device, error) {
	c := &serial.Config{Name: portName, Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
		return Device{}, err
	}
	return Device{port: s}, nil
}

// SetBitrate sets the bitrate of the device.
func (d *Device) SetBitrate(rate int) error {
	str := fmt.Sprintf("S%d\r", rate)
	_, err :=  d.port.Write([]byte(str))
	return err
}

// Open opens the connection to the CAN bus, enabling reception and transmission
// of CAN frames.
func (d *Device) Open() error {
	_, err :=  d.port.Write([]byte("O\r"))
	return err
}

// Close closes the connection to the CAN bus, disabling reception and
// transmission of CAN frames.
func (d *Device) Close() error {
	_, err :=  d.port.Write([]byte("C\r"))
	return err
}

// WriteFrame writes a single Frame to the CAN bus.
func (d *Device) WriteFrame(f Frame) error {
	str := fmt.Sprintf("t%03X%d", f.ID, f.Dlc)
	for i := 0; i < f.Dlc; i++ {
		str = fmt.Sprintf("%s%02X", str, f.Data[i])
	}
	str = str + "\r"
	
	_, err :=  d.port.Write([]byte(str))
	
	return err
}

// ReadFrame reads a single Frame from the CAN bus, blocks until a frame is
// received.
func (d *Device) ReadFrame() (Frame,error) {
	buf := make([]byte, 128)
	_, err := d.port.Read(buf)
	
	if err != nil {
		return Frame{}, err
	}
	
	f := Frame{}
	var dataString string

	fmt.Sscanf(string(buf), "t%3X%1d%s", &f.ID, &f.Dlc, &dataString)
	f.Data = make([]byte, 8)
	for i := 0; i < f.Dlc; i++ {
		fmt.Sscanf(dataString[i*2:i*2+2], "%2X", &f.Data[i])
	}
		
	return f, nil
}
