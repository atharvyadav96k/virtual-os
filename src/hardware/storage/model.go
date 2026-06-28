package storage

import "github.com/atharvyadav96k/virtual-os/hardware/storage/values"

type StoreType int

const (
	Undefined StoreType = iota
	Immediate
	Register
)

type Storage interface {
	GetStorageType() StoreType
	GetStorageValue() values.Value
}
