package aztec

import (
	"github.com/boombuler/barcode/utils"
)

func bitsToWords(stuffedBits *utils.BitList, wordSize int, wordCount int) []int {
	_ = "STUB: not implemented"
	return nil
}

func generateCheckWords(bits *utils.BitList, totalBits, wordSize int) *utils.BitList {
	_ = "STUB: not implemented"
	return nil
}

// bits is guaranteed to be a multiple of the wordSize, so no padding needed

func getGF(wordSize int) *utils.GaloisField { _ = "STUB: not implemented"; return nil }
