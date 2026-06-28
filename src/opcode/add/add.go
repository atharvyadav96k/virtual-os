package add

import (
	"github.com/atharvyadav96k/virtual-os/hardware/storage"
	"github.com/atharvyadav96k/virtual-os/hardware/storage/register"
	"github.com/atharvyadav96k/virtual-os/hardware/storage/values"
	"github.com/atharvyadav96k/virtual-os/opcode"
)

func NewAdd(destIdx int, destination *register.Register, source storage.Storage) Add {
	return Add{
		codeType:    opcode.Add,
		destIdx:     destIdx,
		destination: destination,
		source:      source,
	}
}

func (a *Add) GetCodeType() opcode.CodeType {
	return a.codeType
}

// adds source into destination register, result stored in destination
func (a *Add) Execute() {
	destReg := a.destination.(*register.Register)
	dest := destReg.GetValue()
	src := a.source.GetStorageValue()

	switch dest.Type {
	case values.Integer:
		d, _ := dest.GetInt()
		s, _ := src.GetInt()
		destReg.SetValue(values.NewInt(d + s))
	case values.Float:
		d, _ := dest.GetFloat()
		s, _ := src.GetFloat()
		destReg.SetValue(values.NewFloat(d + s))
	}
}
