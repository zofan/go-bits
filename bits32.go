package bits

type Bits32 uint32

func (b *Bits32) Set(flag Bits32) *Bits32 {
	*b |= flag
	return b
}

func (b *Bits32) Unset(flag Bits32) *Bits32 {
	*b &= ^flag
	return b
}

func (b *Bits32) Toggle(flag Bits32) *Bits32 {
	*b ^= flag
	return b
}

func (b *Bits32) Has(flag Bits32) bool {
	return *b&flag != 0
}
