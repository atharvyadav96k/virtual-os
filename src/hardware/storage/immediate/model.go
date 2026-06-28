package immediate

import (
	"github.com/atharvyadav96k/virtual-os/hardware/storage"
	"github.com/atharvyadav96k/virtual-os/hardware/storage/values"
)

type Immediate struct {
	storageType storage.StoreType
	value       values.Value
}
