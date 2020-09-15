package bits

import (
	"testing"
)

const (
	BitA = 1 << iota
	BitB
	BitC
	BitD
)

func Test8(t *testing.T) {
	b := Bits8(0)

	b.Set(BitA)
	b.Set(BitB)
	b.Set(BitC)
	b.Set(BitD)

	if b != Bits8(15) {
		t.Error(`expected 15, given `, b)
	}

	b.Unset(BitA)
	b.Unset(BitC)

	if b != Bits8(10) {
		t.Error(`expected 10, given `, b)
	}

	b.Toggle(BitC)
	if b != Bits8(14) {
		t.Error(`expected 14, given `, b)
	}

	b.Toggle(BitC)
	if b != Bits8(10) {
		t.Error(`expected 10, given `, b)
	}

	if b.Has(BitA) {
		t.Error(`expected false, given true`)
	}

	if !b.Has(BitB) {
		t.Error(`expected true, given false`)
	}
}

func Test16(t *testing.T) {
	b := Bits16(0)

	b.Set(BitA)
	b.Set(BitB)
	b.Set(BitC)
	b.Set(BitD)

	if b != Bits16(15) {
		t.Error(`expected 15, given `, b)
	}

	b.Unset(BitA)
	b.Unset(BitC)

	if b != Bits16(10) {
		t.Error(`expected 10, given `, b)
	}

	b.Toggle(BitC)
	if b != Bits16(14) {
		t.Error(`expected 14, given `, b)
	}

	b.Toggle(BitC)
	if b != Bits16(10) {
		t.Error(`expected 10, given `, b)
	}

	if b.Has(BitA) {
		t.Error(`expected false, given true`)
	}

	if !b.Has(BitB) {
		t.Error(`expected true, given false`)
	}
}

func Test32(t *testing.T) {
	b := Bits32(0)

	b.Set(BitA)
	b.Set(BitB)
	b.Set(BitC)
	b.Set(BitD)

	if b != Bits32(15) {
		t.Error(`expected 15, given `, b)
	}

	b.Unset(BitA)
	b.Unset(BitC)

	if b != Bits32(10) {
		t.Error(`expected 10, given `, b)
	}

	b.Toggle(BitC)
	if b != Bits32(14) {
		t.Error(`expected 14, given `, b)
	}

	b.Toggle(BitC)
	if b != Bits32(10) {
		t.Error(`expected 10, given `, b)
	}

	if b.Has(BitA) {
		t.Error(`expected false, given true`)
	}

	if !b.Has(BitB) {
		t.Error(`expected true, given false`)
	}
}

func Test64(t *testing.T) {
	b := Bits64(0)

	b.Set(BitA)
	b.Set(BitB)
	b.Set(BitC)
	b.Set(BitD)

	if b != Bits64(15) {
		t.Error(`expected 15, given `, b)
	}

	b.Unset(BitA)
	b.Unset(BitC)

	if b != Bits64(10) {
		t.Error(`expected 10, given `, b)
	}

	b.Toggle(BitC)
	if b != Bits64(14) {
		t.Error(`expected 14, given `, b)
	}

	b.Toggle(BitC)
	if b != Bits64(10) {
		t.Error(`expected 10, given `, b)
	}

	if b.Has(BitA) {
		t.Error(`expected false, given true`)
	}

	if !b.Has(BitB) {
		t.Error(`expected true, given false`)
	}
}
