package cpu

import (
	"github.com/atharvyadav96k/virtual-os/hardware/flags"
	"github.com/atharvyadav96k/virtual-os/hardware/register"
)

type CPU struct {
	general_purpose_register []register.Register
	programCounter           int
	stackPointer             int
	flags                    flags.Flags
}
