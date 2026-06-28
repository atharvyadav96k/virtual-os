package mov

import (
	"github.com/atharvyadav96k/virtual-os/hardware/storage"
	"github.com/atharvyadav96k/virtual-os/hardware/storage/register"
	"github.com/atharvyadav96k/virtual-os/opcode"
)

func NewMov(destination *register.Register, source storage.Storage) Mov {
	return Mov{
		codeType:    opcode.Mov,
		destination: destination,
		source:      source,
	}
}

func (m *Mov) GetCodeType() opcode.CodeType {
	return m.codeType
}

// copies the source value into the destination register
func (m *Mov) Execute() {
	val := m.source.GetStorageValue()
	m.destination.(*register.Register).SetValue(val)
}
