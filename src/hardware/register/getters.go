package register

import "github.com/atharvyadav96k/virtual-os/hardware/values"

// get Register Value
func (r *register) GetValue() values.Value {
	return r.value
}
