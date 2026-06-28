package load

import (
	"github.com/atharvyadav96k/virtual-os/hardware/ram"
	"github.com/atharvyadav96k/virtual-os/hardware/storage/immediate"
	"github.com/atharvyadav96k/virtual-os/hardware/storage/register"
	"github.com/atharvyadav96k/virtual-os/hardware/storage/values"
	"github.com/atharvyadav96k/virtual-os/opcode"
)

func NewLoad(destIdx int, destination *register.Register, address int, r *ram.Ram) Load {
	src := immediate.NewImmediate(values.Integer)
	src.SetValue(values.NewInt(address))
	return Load{
		codeType:    opcode.Load,
		destIdx:     destIdx,
		destination: destination,
		source:      &src,
		ram:         r,
	}
}

func (l *Load) GetCodeType() opcode.CodeType {
	return l.codeType
}

// reads a value from RAM at the address stored in source, writes it to destination register
func (l *Load) Execute() {
	addr, _ := l.source.GetStorageValue().GetInt()
	val := l.ram.Read(addr)
	l.destination.(*register.Register).SetValue(val)
}
