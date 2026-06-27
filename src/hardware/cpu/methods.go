package cpu

import "github.com/atharvyadav96k/virtual-os/hardware/values"

func (c *CPU) SetGeneralPurposeRegisterValue(idx int, value values.Value) {
	c.general_purpose_register[idx].SetValue(value)
}

func (c *CPU) GetGeneralPurposeRegisterValue(idx int) values.Value {
	return c.general_purpose_register[idx].GetValue()
}

func (c *CPU) IncrementProgramCounter() {
	c.programCounter++
}

func (c *CPU) JumpProgramCounter(jump int) {
	c.programCounter += jump
}
