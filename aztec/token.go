package aztec

import (
	"fmt"

	"github.com/boombuler/barcode/utils"
)

type token interface {
	fmt.Stringer
	prev() token
	appendTo(bits *utils.BitList, text []byte)
}

type simpleToken struct {
	token
	value    int
	bitCount byte
}

type binaryShiftToken struct {
	token
	bShiftStart   int
	bShiftByteCnt int
}

func newSimpleToken(prev token, value int, bitCount byte) token {
	_ = "STUB: not implemented"
	return *new(token)
}

func newShiftToken(prev token, bShiftStart int, bShiftCnt int) token {
	_ = "STUB: not implemented"
	return *new(token)
}

func (st *simpleToken) prev() token { _ = "STUB: not implemented"; return *new(token) }

func (st *simpleToken) appendTo(bits *utils.BitList, text []byte) {
	_ = "STUB: not implemented"
	return
}

func (st *simpleToken) String() string { _ = "STUB: not implemented"; return "" }

func (bst *binaryShiftToken) prev() token { _ = "STUB: not implemented"; return *new(token) }

func (bst *binaryShiftToken) appendTo(bits *utils.BitList, text []byte) {
	_ = "STUB: not implemented"
	return
}

// We need a header before the first character, and before
// character 31 when the total byte code is <= 62
// BINARY_SHIFT

// 1 <= binaryShiftByteCode <= 62

// 32 <= binaryShiftCount <= 62 and i == 31

func (bst *binaryShiftToken) String() string { _ = "STUB: not implemented"; return "" }
