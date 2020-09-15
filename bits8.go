package bits

type Bits8 uint8

func (b *Bits8) Set(flag Bits8) {
	*b |= flag
}

func (b *Bits8) Unset(flag Bits8) {
	*b &= ^flag
}

func (b *Bits8) Toggle(flag Bits8) {
	*b ^= flag
}

func (b *Bits8) Has(flag Bits8) bool {
	return *b&flag != 0
}
