// Package code128 can create Code128 barcodes
package code128

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/utils"
)

func strToRunes(str string) []rune { _ = "STUB: not implemented"; return nil }

func shouldUseCTable(nextRunes []rune, curEncoding byte) bool {
	_ = "STUB: not implemented"
	return false
}

func tableContainsRune(table string, r rune) bool { _ = "STUB: not implemented"; return false }

func shouldUseATable(nextRunes []rune, curEncoding byte) bool {
	_ = "STUB: not implemented"
	return false
}

func getCodeIndexList(content []rune) *utils.BitList { _ = "STUB: not implemented"; return nil }

// Encode creates a Code 128 barcode for the given content and color scheme
func EncodeWithColor(content string, color barcode.ColorScheme) (barcode.BarcodeIntCS, error) {
	_ = "STUB: not implemented"
	return *new(barcode.BarcodeIntCS), nil
}

// Encode creates a Code 128 barcode for the given content
func Encode(content string) (barcode.BarcodeIntCS, error) {
	_ = "STUB: not implemented"
	return *new(barcode.BarcodeIntCS), nil
}

func EncodeWithoutChecksum(content string) (barcode.Barcode, error) {
	_ = "STUB: not implemented"
	return *new(barcode.Barcode), nil
}

func EncodeWithoutChecksumWithColor(content string, color barcode.ColorScheme) (barcode.Barcode, error) {
	_ = "STUB: not implemented"
	return *new(barcode.Barcode), nil
}
