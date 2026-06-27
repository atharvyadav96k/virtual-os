package register

import "github.com/atharvyadav96k/virtual-os/hardware/values"

// set value of type which register hold's

func (r *Register) SetValue(value values.Value) {
	r.value = value
}
