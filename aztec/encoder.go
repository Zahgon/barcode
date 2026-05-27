// Package aztec can create Aztec Code barcodes
package aztec

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/utils"
)

const (
	DEFAULT_EC_PERCENT  = 33
	DEFAULT_LAYERS      = 0
	max_nb_bits         = 32
	max_nb_bits_compact = 4
)

var (
	word_size = []int{
		4, 6, 6, 8, 8, 8, 8, 8, 8, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
		12, 12, 12, 12, 12, 12, 12, 12, 12, 12,
	}
)

func totalBitsInLayer(layers int, compact bool) int { _ = "STUB: not implemented"; return 0 }

func stuffBits(bits *utils.BitList, wordSize int) *utils.BitList {
	_ = "STUB: not implemented"
	return nil
}

func generateModeMessage(compact bool, layers, messageSizeInWords int) *utils.BitList {
	_ = "STUB: not implemented"
	return nil
}

func drawModeMessage(matrix *aztecCode, compact bool, matrixSize int, modeMessage *utils.BitList) {
	_ = "STUB: not implemented"
	return
}

func drawBullsEye(matrix *aztecCode, center, size int) { _ = "STUB: not implemented"; return }

// Encode returns an aztec barcode with the given content
func Encode(data []byte, minECCPercent int, userSpecifiedLayers int) (barcode.Barcode, error) {
	_ = "STUB: not implemented"
	return *new(barcode.Barcode), nil
}

// Encode returns an aztec barcode with the given content and color scheme
func EncodeWithColor(data []byte, minECCPercent int, userSpecifiedLayers int, color barcode.ColorScheme) (barcode.Barcode, error) {
	_ = "STUB: not implemented"
	return *new(barcode.Barcode), nil
}

// We look at the possible table sizes in the order Compact1, Compact2, Compact3,
// Compact4, Normal4,...  Normal(i) for i < 4 isn't typically used since Compact(i+1)
// is the same size, but has more data.

// [Re]stuff the bits if this is the first opportunity, or if the
// wordSize has changed

// Compact format only allows 64 data words, though C4 can hold more words than that

// allocate symbol

// no alignment marks in compact mode, alignmentMap is a no-op

// draw data bits

// draw mode message

// draw alignment marks
