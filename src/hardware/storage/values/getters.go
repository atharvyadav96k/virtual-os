package values

import "fmt"

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

func (v Value) GetCharacter() (byte, error) {
	if v.Type != Character {
		return ' ', fmt.Errorf("expected String, got %v", v.Type.String())
	}
	return v.Char, nil
}

func (v Value) GetBool() (bool, error) {
	if v.Type != Bool {
		return false, fmt.Errorf("expected Bool, got %v", v.Type.String())
	}
	return v.Bool, nil
}
