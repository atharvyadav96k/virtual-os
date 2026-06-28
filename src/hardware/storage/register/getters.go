package register

import (
	"github.com/atharvyadav96k/virtual-os/hardware/storage"
	"github.com/atharvyadav96k/virtual-os/hardware/storage/values"
)

// get Register Value
func (r *Register) GetValue() values.Value {
	return r.value
}

func (r *Register) GetStorageType() storage.StoreType {
	return r.storageType
}
