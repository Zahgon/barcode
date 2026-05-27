// Package utils contain some utilities which are needed to create barcodes
package utils

import (
	"image"
	"image/color"

	"github.com/boombuler/barcode"
)

type base1DCode struct {
	*BitList
	kind    string
	content string
	color   barcode.ColorScheme
}

type base1DCodeIntCS struct {
	base1DCode
	checksum int
}

func (c *base1DCode) Content() string { _ = "STUB: not implemented"; return "" }

func (c *base1DCode) Metadata() barcode.Metadata {
	_ = "STUB: not implemented"
	return *new(barcode.Metadata)
}

func (c *base1DCode) ColorModel() color.Model { _ = "STUB: not implemented"; return *new(color.Model) }

func (c *base1DCode) ColorScheme() barcode.ColorScheme {
	_ = "STUB: not implemented"
	return *new(barcode.ColorScheme)
}

func (c *base1DCode) Bounds() image.Rectangle {
	_ = "STUB: not implemented"
	return *new(image.Rectangle)
}

func (c *base1DCode) At(x, y int) color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

func (c *base1DCodeIntCS) CheckSum() int {
	_ = "STUB: not implemented"

	// New1DCodeIntCheckSum creates a new 1D barcode where the bars are represented by the bits in the bars BitList
	return 0
}

func New1DCodeIntCheckSum(codeKind, content string, bars *BitList, checksum int) barcode.BarcodeIntCS {
	_ = "STUB: not implemented"
	return *new(barcode.BarcodeIntCS)
}

// New1DCodeIntCheckSum creates a new 1D barcode where the bars are represented by the bits in the bars BitList
func New1DCodeIntCheckSumWithColor(codeKind, content string, bars *BitList, checksum int, color barcode.ColorScheme) barcode.BarcodeIntCS {
	_ = "STUB: not implemented"
	return *new(barcode.BarcodeIntCS)
}

// New1DCode creates a new 1D barcode where the bars are represented by the bits in the bars BitList
func New1DCode(codeKind, content string, bars *BitList) barcode.Barcode {
	_ = "STUB: not implemented"
	return *new(barcode.Barcode)
}

// New1DCode creates a new 1D barcode where the bars are represented by the bits in the bars BitList
func New1DCodeWithColor(codeKind, content string, bars *BitList, color barcode.ColorScheme) barcode.Barcode {
	_ = "STUB: not implemented"
	return *new(barcode.Barcode)
}
