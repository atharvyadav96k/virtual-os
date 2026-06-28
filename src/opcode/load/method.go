package load

import (
	"github.com/atharvyadav96k/virtual-os/hardware/storage"
	"github.com/atharvyadav96k/virtual-os/hardware/storage/register"
	"github.com/atharvyadav96k/virtual-os/opcode"
)

func (l *Load) GetCodeType() opcode.CodeType {
	return l.codeType
}

func (l *Load) Store(destination register.Register, source storage.Storage) {

}
