package mov

import (
	"github.com/atharvyadav96k/virtual-os/opcode"
)

func NewMovOpCode() Mov {
	return Mov{
		codeType: opcode.Mov,
	}
}
