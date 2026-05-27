package qr

import (
	"github.com/boombuler/barcode/utils"
)

type errorCorrection struct {
	rs *utils.ReedSolomonEncoder
}

var ec = newErrorCorrection()

func newErrorCorrection() *errorCorrection { _ = "STUB: not implemented"; return nil }

func (ec *errorCorrection) calcECC(data []byte, eccCount byte) []byte {
	_ = "STUB: not implemented"
	return nil
}
