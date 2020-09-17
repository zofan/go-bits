package bits

type Bits8 uint8

func (b *Bits8) Set(flag Bits8) *Bits8 {
	*b |= flag
	return b
}

func (b *Bits8) Unset(flag Bits8) *Bits8 {
	*b &= ^flag
	return b
}

func (b *Bits8) Toggle(flag Bits8) *Bits8 {
	*b ^= flag
	return b
}

func (b *Bits8) Has(flag Bits8) bool {
	return *b&flag != 0
}
