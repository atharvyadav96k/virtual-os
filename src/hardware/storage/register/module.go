package register

import (
	"github.com/atharvyadav96k/virtual-os/hardware/storage"
	"github.com/atharvyadav96k/virtual-os/hardware/storage/values"
)

type Register struct {
	storageType storage.StoreType
	// Value inside the register
	value values.Value
}
