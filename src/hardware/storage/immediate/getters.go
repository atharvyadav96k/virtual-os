package immediate

import (
	"github.com/atharvyadav96k/virtual-os/hardware/storage"
	"github.com/atharvyadav96k/virtual-os/hardware/storage/values"
)

func (i *Immediate) GetValue() values.Value {
	return i.value
}

func (i *Immediate) GetStorageType() storage.StoreType {
	return i.storageType
}
