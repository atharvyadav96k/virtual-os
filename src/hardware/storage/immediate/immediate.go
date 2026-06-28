package immediate

import (
	"github.com/atharvyadav96k/virtual-os/hardware/storage"
	"github.com/atharvyadav96k/virtual-os/hardware/storage/values"
)

func NewImmediate(immediateValueType values.ValueType) Immediate {
	return Immediate{
		storageType: storage.Immediate,
		value:       values.NewValue(immediateValueType),
	}
}

func (i *Immediate) SetValue(value values.Value) {
	i.value = value
}

func (i *Immediate) GetValue() values.Value {
	return i.value
}

func (i *Immediate) GetStorageType() storage.StoreType {
	return i.storageType
}

func (i *Immediate) GetStorageValue() values.Value {
	return i.value
}
