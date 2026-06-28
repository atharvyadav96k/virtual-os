package cpu

import (
	"github.com/atharvyadav96k/virtual-os/hardware/cpu/flags"
	"github.com/atharvyadav96k/virtual-os/hardware/storage/register"
	"github.com/atharvyadav96k/virtual-os/opcode"
)

type CPU struct {
	general_purpose_register []register.Register
	programCounter           int
	stackPointer             int
	flags                    flags.Flags
	instruction              opcode.Code
}
