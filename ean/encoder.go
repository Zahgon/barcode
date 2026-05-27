// Package ean can create EAN 8 and EAN 13 barcodes.
package ean

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/utils"
)

type encodedNumber struct {
	LeftOdd  []bool
	LeftEven []bool
	Right    []bool
	CheckSum []bool
}

var encoderTable = map[rune]encodedNumber{
	'0': encodedNumber{
		[]bool{false, false, false, true, true, false, true},
		[]bool{false, true, false, false, true, true, true},
		[]bool{true, true, true, false, false, true, false},
		[]bool{false, false, false, false, false, false},
	},
	'1': encodedNumber{
		[]bool{false, false, true, true, false, false, true},
		[]bool{false, true, true, false, false, true, true},
		[]bool{true, true, false, false, true, true, false},
		[]bool{false, false, true, false, true, true},
	},
	'2': encodedNumber{
		[]bool{false, false, true, false, false, true, true},
		[]bool{false, false, true, true, false, true, true},
		[]bool{true, true, false, true, true, false, false},
		[]bool{false, false, true, true, false, true},
	},
	'3': encodedNumber{
		[]bool{false, true, true, true, true, false, true},
		[]bool{false, true, false, false, false, false, true},
		[]bool{true, false, false, false, false, true, false},
		[]bool{false, false, true, true, true, false},
	},
	'4': encodedNumber{
		[]bool{false, true, false, false, false, true, true},
		[]bool{false, false, true, true, true, false, true},
		[]bool{true, false, true, true, true, false, false},
		[]bool{false, true, false, false, true, true},
	},
	'5': encodedNumber{
		[]bool{false, true, true, false, false, false, true},
		[]bool{false, true, true, true, false, false, true},
		[]bool{true, false, false, true, true, true, false},
		[]bool{false, true, true, false, false, true},
	},
	'6': encodedNumber{
		[]bool{false, true, false, true, true, true, true},
		[]bool{false, false, false, false, true, false, true},
		[]bool{true, false, true, false, false, false, false},
		[]bool{false, true, true, true, false, false},
	},
	'7': encodedNumber{
		[]bool{false, true, true, true, false, true, true},
		[]bool{false, false, true, false, false, false, true},
		[]bool{true, false, false, false, true, false, false},
		[]bool{false, true, false, true, false, true},
	},
	'8': encodedNumber{
		[]bool{false, true, true, false, true, true, true},
		[]bool{false, false, false, true, false, false, true},
		[]bool{true, false, false, true, false, false, false},
		[]bool{false, true, false, true, true, false},
	},
	'9': encodedNumber{
		[]bool{false, false, false, true, false, true, true},
		[]bool{false, false, true, false, true, true, true},
		[]bool{true, true, true, false, true, false, false},
		[]bool{false, true, true, false, true, false},
	},
}

func calcCheckNum(code string) rune { _ = "STUB: not implemented"; return 0 }

func encodeEAN8(code string) *utils.BitList { _ = "STUB: not implemented"; return nil }

func encodeEAN13(code string) *utils.BitList { _ = "STUB: not implemented"; return nil }

// Left

// Encode returns a EAN 8 or EAN 13 barcode for the given code and color scheme
func EncodeWithColor(code string, color barcode.ColorScheme) (barcode.BarcodeIntCS, error) {
	_ = "STUB: not implemented"
	return *new(barcode.BarcodeIntCS), nil
}

// Encode returns a EAN 8 or EAN 13 barcode for the given code
func Encode(code string) (barcode.BarcodeIntCS, error) {
	_ = "STUB: not implemented"
	return *new(barcode.BarcodeIntCS), nil
}
