package bits

type Bits32 uint8

func (b *Bits32) Set(flag Bits32) {
	*b |= flag
}

func (b *Bits32) Unset(flag Bits32) {
	*b &= ^flag
}

func (b *Bits32) Toggle(flag Bits32) {
	*b ^= flag
}

func (b *Bits32) Has(flag Bits32) bool {
	return *b&flag != 0
}
