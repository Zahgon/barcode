package barcode

import (
	"image"
	"image/color"
)

type wrapFunc func(x, y int) color.Color

type scaledBarcode struct {
	wrapped     Barcode
	wrapperFunc wrapFunc
	rect        image.Rectangle
}

type intCSscaledBC struct {
	scaledBarcode
}

func (bc *scaledBarcode) Content() string { _ = "STUB: not implemented"; return "" }

func (bc *scaledBarcode) Metadata() Metadata { _ = "STUB: not implemented"; return *new(Metadata) }

func (bc *scaledBarcode) ColorModel() color.Model {
	_ = "STUB: not implemented"
	return *new(color.Model)
}

func (bc *scaledBarcode) Bounds() image.Rectangle {
	_ = "STUB: not implemented"
	return *new(image.Rectangle)
}

func (bc *scaledBarcode) At(x, y int) color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

func (bc *intCSscaledBC) CheckSum() int { _ = "STUB: not implemented"; return 0 }

// Scale returns a resized barcode with the given width and height.
func Scale(bc Barcode, width, height int) (Barcode, error) {
	_ = "STUB: not implemented"
	return *new(Barcode), nil
}

// Scale returns a resized barcode with the given width, height and fill color.
func ScaleWithFill(bc Barcode, width, height int, fill color.Color) (Barcode, error) {
	_ = "STUB: not implemented"
	return *new(Barcode), nil
}

func newScaledBC(wrapped Barcode, wrapperFunc wrapFunc, rect image.Rectangle) Barcode {
	_ = "STUB: not implemented"
	return *new(Barcode)
}

func scale2DCode(bc Barcode, width, height int, fill color.Color) (Barcode, error) {
	_ = "STUB: not implemented"
	return *new(Barcode), nil
}

func scale1DCode(bc Barcode, width, height int, fill color.Color) (Barcode, error) {
	_ = "STUB: not implemented"
	return *new(Barcode), nil
}
