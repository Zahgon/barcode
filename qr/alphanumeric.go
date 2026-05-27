package qr

import (
	"github.com/boombuler/barcode/utils"
)

const charSet string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ $%*+-./:"

func stringToAlphaIdx(content string) <-chan int { _ = "STUB: not implemented"; return nil }

func encodeAlphaNumeric(content string, ecl ErrorCorrectionLevel) (*utils.BitList, *versionInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
