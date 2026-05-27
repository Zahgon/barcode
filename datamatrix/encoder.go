// Package datamatrix can create Datamatrix barcodes
package datamatrix

import (
	"github.com/boombuler/barcode"
)

// FNC1 is the codeword for the Function 1 Symbol Character to
// differentiate a GS1 DataMatrix from other Data Matrix symbols.
//
// It is used as both a start character and a separator of GS1 element
// strings.
const FNC1 byte = 232

// Encode returns a Datamatrix barcode for the given content and color scheme
func EncodeWithColor(content string, color barcode.ColorScheme) (barcode.Barcode, error) {
	_ = "STUB: not implemented"
	return *new(barcode.Barcode), nil
}

// Encode returns a Datamatrix barcode for the given content
func Encode(content string) (barcode.Barcode, error) {
	_ = "STUB: not implemented"
	return *new(barcode.Barcode), nil
}

func render(data []byte, size *dmCodeSize, color barcode.ColorScheme) *datamatrixCode {
	_ = "STUB: not implemented"
	return nil
}

func encodeText(content string) []byte { _ = "STUB: not implemented"; return nil }

// two numbers...

// not correct... needs to be redone later...

func addPadding(data []byte, toCount int) []byte { _ = "STUB: not implemented"; return nil }
