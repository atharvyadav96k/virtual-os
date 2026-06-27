package values

type ValueType uint8

const (
	Null ValueType = iota
	Integer
	Float
	String
	Bool
)

type Value struct {
	Type  ValueType
	Int   int
	Float float32
	Str   string
	Bool  bool
}
