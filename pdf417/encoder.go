// Package pdf417 can create PDF-417 barcodes
package pdf417

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/utils"
)

const (
	padding_codeword = 900
)

// Encodes the given data and color scheme as PDF417 barcode.
// securityLevel should be between 0 and 8. The higher the number, the more
// additional error-correction codes are added.
func EncodeWithColor(data string, securityLevel byte, color barcode.ColorScheme) (barcode.Barcode, error) {
	_ = "STUB: not implemented"
	return *new(barcode.Barcode), nil
}

// Encodes the given data as PDF417 barcode.
// securityLevel should be between 0 and 8. The higher the number, the more
// additional error-correction codes are added.
func Encode(data string, securityLevel byte) (barcode.Barcode, error) {
	_ = "STUB: not implemented"
	return *new(barcode.Barcode), nil
}

func encodeData(dataWords []int, columns int, sl securitylevel) ([]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getLeftCodeWord(rowNum int, rows int, columns int, securityLevel byte) int {
	_ = "STUB: not implemented"
	return 0
}

func getRightCodeWord(rowNum int, rows int, columns int, securityLevel byte) int {
	_ = "STUB: not implemented"
	return 0
}

func min(a, b int) int { _ = "STUB: not implemented"; return 0 }

func getPadding(dataCount int, ecCount int, columns int) []int {
	_ = "STUB: not implemented"
	return nil
}

func renderBarcode(codes [][]int) *utils.BitList { _ = "STUB: not implemented"; return nil }
