package mov

import (
	"github.com/atharvyadav96k/virtual-os/hardware/storage"
	"github.com/atharvyadav96k/virtual-os/hardware/storage/register"
	"github.com/atharvyadav96k/virtual-os/opcode"
)

func (m *Mov) GetCodeType() opcode.CodeType {
	return m.codeType
}

func (n *Mov) Store(destination register.Register, source storage.Storage) {

}
