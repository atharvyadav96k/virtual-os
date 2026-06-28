package cpu

import (
	"github.com/atharvyadav96k/virtual-os/hardware/flags"
	"github.com/atharvyadav96k/virtual-os/hardware/storage/register"
)

func NewCpu() CPU {
	return CPU{
		general_purpose_register: register.NewRegisters(7),
		programCounter:           0,
		stackPointer:             0,
		flags:                    flags.NewFlags(),
	}
}
