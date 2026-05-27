package pdf417

import (
	"image"
	"image/color"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/utils"
)

type pdfBarcode struct {
	data  string
	width int
	code  *utils.BitList
	color barcode.ColorScheme
}

func (c *pdfBarcode) Metadata() barcode.Metadata {
	_ = "STUB: not implemented"
	return *new(barcode.Metadata)
}

func (c *pdfBarcode) Content() string { _ = "STUB: not implemented"; return "" }

func (c *pdfBarcode) ColorModel() color.Model { _ = "STUB: not implemented"; return *new(color.Model) }

func (c *pdfBarcode) ColorScheme() barcode.ColorScheme {
	_ = "STUB: not implemented"
	return *new(barcode.ColorScheme)
}

func (c *pdfBarcode) Bounds() image.Rectangle {
	_ = "STUB: not implemented"
	return *new(image.Rectangle)
}

func (c *pdfBarcode) At(x, y int) color.Color { _ = "STUB: not implemented"; return *new(color.Color) }
