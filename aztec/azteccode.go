package aztec

import (
	"image"
	"image/color"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/utils"
)

type aztecCode struct {
	*utils.BitList
	size    int
	content []byte
	color   barcode.ColorScheme
}

func newAztecCode(size int, color barcode.ColorScheme) *aztecCode {
	_ = "STUB: not implemented"
	return nil
}

func (c *aztecCode) Content() string { _ = "STUB: not implemented"; return "" }

func (c *aztecCode) Metadata() barcode.Metadata {
	_ = "STUB: not implemented"
	return *new(barcode.Metadata)
}

func (c *aztecCode) ColorModel() color.Model { _ = "STUB: not implemented"; return *new(color.Model) }

func (c *aztecCode) ColorScheme() barcode.ColorScheme {
	_ = "STUB: not implemented"
	return *new(barcode.ColorScheme)
}

func (c *aztecCode) Bounds() image.Rectangle {
	_ = "STUB: not implemented"
	return *new(image.Rectangle)
}

func (c *aztecCode) At(x, y int) color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

func (c *aztecCode) set(x, y int) { _ = "STUB: not implemented"; return }

func (c *aztecCode) string() string { _ = "STUB: not implemented"; return "" }
