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

func NewString(v string) Value {
	return Value{
		Type: String,
		Str:  v,
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
	case String:
		return NewString("")
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
