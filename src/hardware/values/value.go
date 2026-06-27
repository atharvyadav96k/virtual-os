package values

import (
	"fmt"
)

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

func (v Value) GetInt() (int, error) {
	if v.Type != Integer {
		return 0, fmt.Errorf("expected Integer, got %v", v.Type.String())
	}
	return v.Int, nil
}

func (v Value) GetFloat() (float32, error) {
	if v.Type != Float {
		return 0, fmt.Errorf("expected Float, got %v", v.Type.String())
	}
	return v.Float, nil
}

func (v Value) GetString() (string, error) {
	if v.Type != String {
		return "", fmt.Errorf("expected String, got %v", v.Type.String())
	}
	return v.Str, nil
}

func (v Value) GetBool() (bool, error) {
	if v.Type != Bool {
		return false, fmt.Errorf("expected Bool, got %v", v.Type.String())
	}
	return v.Bool, nil
}
