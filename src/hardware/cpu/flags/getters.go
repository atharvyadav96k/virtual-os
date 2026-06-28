package flags

func (f *Flags) IsZero() bool {
	return f.zero
}

func (f *Flags) IsNegative() bool {
	return f.negative
}

func (f *Flags) IsOverflow() bool {
	return f.overFlow
}

func (f *Flags) IsKernel() bool {
	return f.kernel
}
