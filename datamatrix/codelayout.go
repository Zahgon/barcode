package datamatrix

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/utils"
)

type codeLayout struct {
	matrix *utils.BitList
	occupy *utils.BitList
	size   *dmCodeSize
	color  barcode.ColorScheme
}

func newCodeLayout(size *dmCodeSize, color barcode.ColorScheme) *codeLayout {
	_ = "STUB: not implemented"
	return nil
}

func (l *codeLayout) Occupied(row, col int) bool { _ = "STUB: not implemented"; return false }

func (l *codeLayout) Set(row, col int, value, bitNum byte) { _ = "STUB: not implemented"; return }

func (l *codeLayout) SetSimple(row, col int, value byte) { _ = "STUB: not implemented"; return }

func (l *codeLayout) Corner1(value byte) { _ = "STUB: not implemented"; return }

func (l *codeLayout) Corner2(value byte) { _ = "STUB: not implemented"; return }

func (l *codeLayout) Corner3(value byte) { _ = "STUB: not implemented"; return }

func (l *codeLayout) Corner4(value byte) { _ = "STUB: not implemented"; return }

func (l *codeLayout) SetValues(data []byte) { _ = "STUB: not implemented"; return }

func (l *codeLayout) Merge() *datamatrixCode { _ = "STUB: not implemented"; return nil }

//dotted horizontal lines

//solid horizontal line

//dotted vertical lines

//solid vertical line
