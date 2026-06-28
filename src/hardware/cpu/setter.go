package cpu

import (
	"fmt"

	"github.com/atharvyadav96k/virtual-os/hardware/storage/values"
)

func (c *CPU) SetGeneralPurposeRegisterValue(idx int, value values.Value) error {
	if len(c.general_purpose_register) >= idx {
		return fmt.Errorf("Invalid operation: General Purpose register don't exists")
	}
	c.general_purpose_register[idx].SetValue(value)
	return nil
}
