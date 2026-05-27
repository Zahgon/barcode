package pdf417

type encodingMode byte

type subMode byte

const (
	encText encodingMode = iota
	encNumeric
	encBinary

	subUpper subMode = iota
	subLower
	subMixed
	subPunct

	latch_to_text        = 900
	latch_to_byte_padded = 901
	latch_to_numeric     = 902
	latch_to_byte        = 924
	shift_to_byte        = 913

	min_numeric_count = 13
)

var (
	mixedMap map[rune]int
	punctMap map[rune]int
)

func init() {
	mixedMap = make(map[rune]int)
	mixedRaw := []rune{
		48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 38, 13, 9, 44, 58,
		35, 45, 46, 36, 47, 43, 37, 42, 61, 94, 0, 32, 0, 0, 0,
	}
	for idx, ch := range mixedRaw {
		if ch > 0 {
			mixedMap[ch] = idx
		}
	}

	punctMap = make(map[rune]int)
	punctRaw := []rune{
		59, 60, 62, 64, 91, 92, 93, 95, 96, 126, 33, 13, 9, 44, 58,
		10, 45, 46, 36, 47, 34, 124, 42, 40, 41, 63, 123, 125, 39, 0,
	}
	for idx, ch := range punctRaw {
		if ch > 0 {
			punctMap[ch] = idx
		}
	}
}

func determineConsecutiveDigitCount(data []rune) int { _ = "STUB: not implemented"; return 0 }

func encodeNumeric(digits []rune) ([]int, error) { _ = "STUB: not implemented"; return nil, nil }

func determineConsecutiveTextCount(msg []rune) int { _ = "STUB: not implemented"; return 0 }

func encodeText(text []rune, submode subMode) (subMode, []int) {
	_ = "STUB: not implemented"
	return *new(subMode), nil
}

//space

// lower latch

// mixed latch

// punctuation switch

//space

//upper switch

//mixed latch

//punctuation switch

//upper latch

//lower latch

//punctuation latch

//punctuation switch

//subPunct

//upper latch

func determineConsecutiveBinaryCount(msg []byte) int { _ = "STUB: not implemented"; return 0 }

func encodeBinary(data []byte, startmode encodingMode) []int { _ = "STUB: not implemented"; return nil }

// Encode sixpacks

//Encode rest (remaining n<5 bytes if any)

func highlevelEncode(dataStr string) ([]int, error) { _ = "STUB: not implemented"; return nil, nil }
