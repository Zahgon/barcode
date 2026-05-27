// Package twooffive can create interleaved and standard "2 of 5" barcodes.
package twooffive

import (
	"github.com/boombuler/barcode"
)

const patternWidth = 5

type pattern [patternWidth]bool
type encodeInfo struct {
	start  []bool
	end    []bool
	widths map[bool]int
}

var (
	encodingTable = map[rune]pattern{
		'0': pattern{false, false, true, true, false},
		'1': pattern{true, false, false, false, true},
		'2': pattern{false, true, false, false, true},
		'3': pattern{true, true, false, false, false},
		'4': pattern{false, false, true, false, true},
		'5': pattern{true, false, true, false, false},
		'6': pattern{false, true, true, false, false},
		'7': pattern{false, false, false, true, true},
		'8': pattern{true, false, false, true, false},
		'9': pattern{false, true, false, true, false},
	}

	modes = map[bool]encodeInfo{
		false: encodeInfo{ // non-interleaved
			start: []bool{true, true, false, true, true, false, true, false},
			end:   []bool{true, true, false, true, false, true, true},
			widths: map[bool]int{
				true:  3,
				false: 1,
			},
		},
		true: encodeInfo{ // interleaved
			start: []bool{true, false, true, false},
			end:   []bool{true, true, true, false, true},
			widths: map[bool]int{
				true:  3,
				false: 1,
			},
		},
	}
	nonInterleavedSpace = pattern{false, false, false, false, false}
)

// AddCheckSum calculates the correct check-digit and appends it to the given content.
func AddCheckSum(content string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Encode creates a codabar barcode for the given content and color scheme
func EncodeWithColor(content string, interleaved bool, color barcode.ColorScheme) (barcode.Barcode, error) {
	_ = "STUB: not implemented"
	return *new(barcode.Barcode), nil
}

// Encode creates a codabar barcode for the given content
func Encode(content string, interleaved bool) (barcode.Barcode, error) {
	_ = "STUB: not implemented"
	return *new(barcode.Barcode), nil
}
