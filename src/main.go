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
	computer.hardware.LoadProgram(0, []string{
		"MOV R0 #42",
		"ADD R0 #8",
		"HALT",
	})
	if err := computer.hardware.Run(); err != nil {
		panic(err)
	}
}
