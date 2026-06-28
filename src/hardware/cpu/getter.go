package cpu

import "github.com/atharvyadav96k/virtual-os/hardware/storage/values"

func (c *CPU) GetGeneralPurposeRegisterValue(idx int) values.Value {
	if len(c.general_purpose_register) >= idx {
		return values.NewNull()
	}
	return c.general_purpose_register[idx].GetValue()
}
