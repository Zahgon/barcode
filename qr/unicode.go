package qr

import (
	"github.com/boombuler/barcode/utils"
)

func encodeUnicode(content string, ecl ErrorCorrectionLevel) (*utils.BitList, *versionInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// It's not correct to add the unicode bytes to the result directly but most readers can't handle the
// required ECI header...
