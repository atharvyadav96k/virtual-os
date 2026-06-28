package load

import "github.com/atharvyadav96k/virtual-os/opcode"

func NewLoadOpCode() Load {
	return Load{
		codeType: opcode.Load,
	}
}
