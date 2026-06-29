package values

func NewNull() Value {
	return Value{Type: Null}
}

func NewInt(v int) Value {
	return Value{
		Type: Integer,
		Int:  v,
	}
}

func NewFloat(v float32) Value {
	return Value{
		Type:  Float,
		Float: v,
	}
}

func NewCharacter(v byte) Value {
	return Value{
		Type: Character,
		Char: v,
	}
}

func NewBool(v bool) Value {
	return Value{
		Type: Bool,
		Bool: v,
	}
}

func NewValue(valueType ValueType) Value {
	switch valueType {
	case Character:
		return NewCharacter(' ')
	case Float:
		return NewFloat(0)
	case Integer:
		return NewInt(0)
	case Bool:
		return NewBool(false)
	default:
		return NewNull()
	}
}
