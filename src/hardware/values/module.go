package values

type ValueType uint8

const (
	Null ValueType = iota
	Integer
	Float
	String
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
	case String:
		return "String"
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
	Str   string
	Bool  bool
}
