package immediate

import "github.com/atharvyadav96k/virtual-os/hardware/storage/values"

func NewImmediate(immediateValueType values.ValueType) Immediate {
	return Immediate{
		value: values.NewValue(immediateValueType),
	}
}
