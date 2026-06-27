package flags

func (f *Flags) SetZero(flag bool) {
	f.zero = flag
}

func (f *Flags) SetNegative(flag bool) {
	f.negative = flag
}

func (f *Flags) SetOverflow(flag bool) {
	f.overFlow = flag
}

func (f *Flags) SetKernel(flag bool) {
	f.kernel = flag
}
