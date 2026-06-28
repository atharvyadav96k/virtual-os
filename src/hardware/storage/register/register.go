package register

import (
	"github.com/atharvyadav96k/virtual-os/hardware/storage/values"
)

func NewRegister(registerValueType values.ValueType) Register {
	return Register{
		value: values.NewValue(registerValueType),
	}
}

func NewRegisters(count int) []Register {
	registers := make([]Register, count)
	for i := range count {
		registers[i] = NewRegister(values.Null)
	}
	return registers
}
