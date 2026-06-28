package load

import (
	"github.com/atharvyadav96k/virtual-os/hardware/ram"
	"github.com/atharvyadav96k/virtual-os/hardware/storage"
	"github.com/atharvyadav96k/virtual-os/opcode"
)

type Load struct {
	codeType    opcode.CodeType
	destIdx     int
	destination storage.Storage
	source      storage.Storage
	ram         *ram.Ram
}
