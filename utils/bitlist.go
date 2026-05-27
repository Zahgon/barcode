package utils

// BitList is a list that contains bits
type BitList struct {
	count int
	data  []int32
}

// NewBitList returns a new BitList with the given length
// all bits are initialize with false
func NewBitList(capacity int) *BitList { _ = "STUB: not implemented"; return nil }

// Len returns the number of contained bits
func (bl *BitList) Len() int { _ = "STUB: not implemented"; return 0 }

func (bl *BitList) grow() { _ = "STUB: not implemented"; return }

// AddBit appends the given bits to the end of the list
func (bl *BitList) AddBit(bits ...bool) { _ = "STUB: not implemented"; return }

// SetBit sets the bit at the given index to the given value
func (bl *BitList) SetBit(index int, value bool) { _ = "STUB: not implemented"; return }

// GetBit returns the bit at the given index
func (bl *BitList) GetBit(index int) bool { _ = "STUB: not implemented"; return false }

// AddByte appends all 8 bits of the given byte to the end of the list
func (bl *BitList) AddByte(b byte) { _ = "STUB: not implemented"; return }

// AddBits appends the last (LSB) 'count' bits of 'b' the the end of the list
func (bl *BitList) AddBits(b int, count byte) { _ = "STUB: not implemented"; return }

// GetBytes returns all bits of the BitList as a []byte
func (bl *BitList) GetBytes() []byte { _ = "STUB: not implemented"; return nil }

// IterateBytes iterates through all bytes contained in the BitList
func (bl *BitList) IterateBytes() <-chan byte { _ = "STUB: not implemented"; return nil }
