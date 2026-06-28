package opcode

type CodeType int

const (
	Mov CodeType = iota
	Load
	Add
	Halt
)

type Code interface {
	GetCodeType() CodeType
	Execute()
}
