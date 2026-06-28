package opcode

type CodeType int

const (
	Mov CodeType = iota
	Load
)

type Code interface {
	GetCodeType() CodeType
}
