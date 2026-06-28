package load

import (
	"github.com/atharvyadav96k/virtual-os/hardware/storage"
	"github.com/atharvyadav96k/virtual-os/opcode"
)

type Load struct {
	codeType opcode.CodeType
	// always register
	destination storage.Storage
	// ram
	source storage.Storage
}
