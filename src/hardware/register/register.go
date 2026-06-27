package register

import "github.com/atharvyadav96k/virtual-os/hardware/values"

func NewRegister(registerValueType values.ValueType) register {
	switch registerValueType {
	case values.String:
		return register{
			value: values.NewFloat(0),
		}
	case values.Integer:
		return register{
			value: values.NewInt(0),
		}
	case values.Bool:
		return register{
			value: values.NewBool(false),
		}
	default:
		return register{
			value: values.NewNull(),
		}
	}
}
