package bits

type Bits64 uint8

func (b *Bits64) Set(flag Bits64) {
	*b |= flag
}

func (b *Bits64) Unset(flag Bits64) {
	*b &= ^flag
}

func (b *Bits64) Toggle(flag Bits64) {
	*b ^= flag
}

func (b *Bits64) Has(flag Bits64) bool {
	return *b&flag != 0
}
