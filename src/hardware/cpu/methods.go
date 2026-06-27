package cpu

import (
	"fmt"

	"github.com/atharvyadav96k/virtual-os/hardware/values"
)

func (c *CPU) SetGeneralPurposeRegisterValue(idx int, value values.Value) error {
	if len(c.general_purpose_register) >= idx {
		return fmt.Errorf("Invalid operation: General Purpose register don't exists")
	}
	c.general_purpose_register[idx].SetValue(value)
	return nil
}

func (c *CPU) GetGeneralPurposeRegisterValue(idx int) values.Value {
	if len(c.general_purpose_register) >= idx {
		return values.NewNull()
	}
	return c.general_purpose_register[idx].GetValue()
}

func (c *CPU) IncrementProgramCounter() {
	c.programCounter++
}

func (c *CPU) JumpProgramCounter(jump int) {
	c.programCounter += jump
}
