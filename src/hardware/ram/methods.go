package ram

import "github.com/atharvyadav96k/virtual-os/hardware/storage/values"

func (r *Ram) Read(address int) values.Value {
	if address < 0 || address >= len(r.storage) {
		return values.NewNull()
	}
	return r.storage[address]
}

func (r *Ram) Write(address int, val values.Value) {
	if address < 0 || address >= len(r.storage) {
		return
	}
	r.storage[address] = val
}
