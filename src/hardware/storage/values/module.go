package values

type ValueType uint8

const (
	Null ValueType = iota
	Integer
	Float
	Character
	Bool
)

func (v ValueType) String() string {
	switch v {
	case Null:
		return "Null"
	case Integer:
		return "Integer"
	case Float:
		return "Float"
	case Character:
		return "Character"
	case Bool:
		return "Bool"
	default:
		return "Unknown"
	}
}

type Value struct {
	Type  ValueType
	Int   int
	Float float32
	Char  byte
	Bool  bool
}
