package ram

import "github.com/atharvyadav96k/virtual-os/hardware/storage/values"

func NewRam(size int) Ram {
	return Ram{
		storage: make([]values.Value, size),
	}
}

func (r *Ram) Size() int {
	return len(r.storage)
}
