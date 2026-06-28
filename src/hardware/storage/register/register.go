package register

import (
	"github.com/atharvyadav96k/virtual-os/hardware/storage"
	"github.com/atharvyadav96k/virtual-os/hardware/storage/values"
)

func NewRegister(registerValueType values.ValueType) Register {
	return Register{
		storageType: storage.Register,
		value:       values.NewValue(registerValueType),
	}
}

func NewRegisters(count int) []Register {
	registers := make([]Register, count)
	for i := range count {
		registers[i] = NewRegister(values.Null)
	}
	return registers
}

func (r *Register) SetValue(value values.Value) {
	r.value = value
}

// get Register Value
func (r *Register) GetValue() values.Value {
	return r.value
}

func (r *Register) GetStorageType() storage.StoreType {
	return r.storageType
}

func (r *Register) GetStorageValue() values.Value {
	return r.value
}
