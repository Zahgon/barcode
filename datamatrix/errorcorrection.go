package datamatrix

import (
	"github.com/boombuler/barcode/utils"
)

type errorCorrection struct {
	rs *utils.ReedSolomonEncoder
}

var ec *errorCorrection = newErrorCorrection()

func newErrorCorrection() *errorCorrection { _ = "STUB: not implemented"; return nil }

func (ec *errorCorrection) calcECC(data []byte, size *dmCodeSize) []byte {
	_ = "STUB: not implemented"
	return nil

	// make some space for error correction codes
}

// copy the data for the current block to buff

// calc the error correction codes

// and append them to the result
