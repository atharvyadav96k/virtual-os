package opcode

type CodeType int

const (
	Mov CodeType = iota
	Load
	Add
)

type Code interface {
	GetCodeType() CodeType
}
