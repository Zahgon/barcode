package qr

import (
	"image"
	"image/color"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/utils"
)

type qrcode struct {
	dimension int
	data      *utils.BitList
	content   string
	color     barcode.ColorScheme
}

func (qr *qrcode) Content() string { _ = "STUB: not implemented"; return "" }

func (qr *qrcode) Metadata() barcode.Metadata {
	_ = "STUB: not implemented"
	return *new(barcode.Metadata)
}

func (qr *qrcode) ColorModel() color.Model { _ = "STUB: not implemented"; return *new(color.Model) }

func (c *qrcode) ColorScheme() barcode.ColorScheme {
	_ = "STUB: not implemented"
	return *new(barcode.ColorScheme)
}

func (qr *qrcode) Bounds() image.Rectangle { _ = "STUB: not implemented"; return *new(image.Rectangle) }

func (qr *qrcode) At(x, y int) color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

func (qr *qrcode) Get(x, y int) bool { _ = "STUB: not implemented"; return false }

func (qr *qrcode) Set(x, y int, val bool) { _ = "STUB: not implemented"; return }

func (qr *qrcode) calcPenalty() uint { _ = "STUB: not implemented"; return 0 }

func (qr *qrcode) calcPenaltyRule1() uint { _ = "STUB: not implemented"; return 0 }

func (qr *qrcode) calcPenaltyRule2() uint { _ = "STUB: not implemented"; return 0 }

func (qr *qrcode) calcPenaltyRule3() uint { _ = "STUB: not implemented"; return 0 }

func (qr *qrcode) calcPenaltyRule4() uint { _ = "STUB: not implemented"; return 0 }

func newBarCodeWithColor(dim int, color barcode.ColorScheme) *qrcode {
	_ = "STUB: not implemented"
	return nil
}

func newBarcode(dim int) *qrcode { _ = "STUB: not implemented"; return nil }
