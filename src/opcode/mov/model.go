package mov

import (
	"github.com/atharvyadav96k/virtual-os/hardware/storage"
	"github.com/atharvyadav96k/virtual-os/opcode"
)

type Mov struct {
	codeType    opcode.CodeType
	destIdx     int
	destination storage.Storage
	source      storage.Storage
}
