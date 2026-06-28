package mov

import (
	"github.com/atharvyadav96k/virtual-os/hardware/storage"
	"github.com/atharvyadav96k/virtual-os/opcode"
)

type Mov struct {
	codeType opcode.CodeType
	// always register
	destination storage.Storage
	// can be register or immediate
	source storage.Storage
}
