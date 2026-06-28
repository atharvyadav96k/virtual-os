package cpu

import (
	"fmt"

	"github.com/atharvyadav96k/virtual-os/hardware/cpu/flags"
	"github.com/atharvyadav96k/virtual-os/hardware/storage/register"
	"github.com/atharvyadav96k/virtual-os/hardware/storage/values"
)

func NewCpu() CPU {
	return CPU{
		general_purpose_register: register.NewRegisters(7),
		programCounter:           0,
		stackPointer:             0,
		flags:                    flags.NewFlags(),
	}
}

func (c *CPU) SetGeneralPurposeRegisterValue(idx int, value values.Value) error {
	if len(c.general_purpose_register) >= idx {
		return fmt.Errorf("Invalid operation: General Purpose register don't exists")
	}
	c.general_purpose_register[idx].SetValue(value)
	return nil
}

func (c *CPU) IncrementProgramCounter() {
	c.programCounter++
}

func (c *CPU) JumpProgramCounter(jump int) {
	c.programCounter += jump
}

func (c *CPU) GetGeneralPurposeRegisterValue(idx int) values.Value {
	if len(c.general_purpose_register) >= idx {
		return values.NewNull()
	}
	return c.general_purpose_register[idx].GetValue()
}
