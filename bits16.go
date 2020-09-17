package bits

type Bits16 uint16

func (b *Bits16) Set(flag Bits16) *Bits16 {
	*b |= flag
	return b
}

func (b *Bits16) Unset(flag Bits16) *Bits16 {
	*b &= ^flag
	return b
}

func (b *Bits16) Toggle(flag Bits16) *Bits16 {
	*b ^= flag
	return b
}

func (b *Bits16) Has(flag Bits16) bool {
	return *b&flag != 0
}
