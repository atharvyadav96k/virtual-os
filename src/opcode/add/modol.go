package add

import (
	"github.com/atharvyadav96k/virtual-os/hardware/storage"
	"github.com/atharvyadav96k/virtual-os/opcode"
)

type Add struct {
	codeType    opcode.CodeType
	destIdx     int
	destination storage.Storage
	source      storage.Storage
}
