package main

import "github.com/atharvyadav96k/virtual-os/hardware"

type computer struct {
	hardware hardware.Hardware
}

func NewComputer() computer {
	return computer{
		hardware: hardware.NewHardware(),
	}
}

func main() {
	computer := NewComputer()
	computer.hardware.
}
