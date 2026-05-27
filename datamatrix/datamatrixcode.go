package datamatrix

import (
	"image"
	"image/color"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/utils"
)

type datamatrixCode struct {
	*utils.BitList
	*dmCodeSize
	content string
	color   barcode.ColorScheme
}

func newDataMatrixCodeWithColor(size *dmCodeSize, color barcode.ColorScheme) *datamatrixCode {
	_ = "STUB: not implemented"
	return nil
}

func newDataMatrixCode(size *dmCodeSize) *datamatrixCode { _ = "STUB: not implemented"; return nil }

func (c *datamatrixCode) Content() string { _ = "STUB: not implemented"; return "" }

func (c *datamatrixCode) Metadata() barcode.Metadata {
	_ = "STUB: not implemented"
	return *new(barcode.Metadata)
}

func (c *datamatrixCode) ColorModel() color.Model {
	_ = "STUB: not implemented"
	return *new(color.Model)
}

func (c *datamatrixCode) ColorScheme() barcode.ColorScheme {
	_ = "STUB: not implemented"
	return *new(barcode.ColorScheme)
}

func (c *datamatrixCode) Bounds() image.Rectangle {
	_ = "STUB: not implemented"
	return *new(image.Rectangle)
}

func (c *datamatrixCode) At(x, y int) color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

func (c *datamatrixCode) get(x, y int) bool { _ = "STUB: not implemented"; return false }

func (c *datamatrixCode) set(x, y int, value bool) { _ = "STUB: not implemented"; return }
