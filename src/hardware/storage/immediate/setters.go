package immediate

import "github.com/atharvyadav96k/virtual-os/hardware/storage/values"

func (i *Immediate) SetValue(value values.Value) {
	i.value = value
}
