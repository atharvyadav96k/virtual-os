package register

import (
	"github.com/atharvyadav96k/virtual-os/hardware/values"
)

func NewRegister(registerValueType values.ValueType) Register {
	switch registerValueType {
	case values.String:
		return Register{
			value: values.NewString(""),
		}
	case values.Float:
		return Register{
			value: values.NewFloat(0),
		}
	case values.Integer:
		return Register{
			value: values.NewInt(0),
		}
	case values.Bool:
		return Register{
			value: values.NewBool(false),
		}
	default:
		return Register{
			value: values.NewNull(),
		}
	}
}

func NewRegisters(count int) []Register {
	registers := make([]Register, count)
	for i := range count {
		registers[i] = NewRegister(values.Null)
	}
	return registers
}
