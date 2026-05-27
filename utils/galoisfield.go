package utils

// GaloisField encapsulates galois field arithmetics
type GaloisField struct {
	Size    int
	Base    int
	ALogTbl []int
	LogTbl  []int
}

// NewGaloisField creates a new galois field
func NewGaloisField(pp, fieldSize, b int) *GaloisField { _ = "STUB: not implemented"; return nil }

func (gf *GaloisField) Zero() *GFPoly { _ = "STUB: not implemented"; return nil }

// AddOrSub add or substract two numbers
func (gf *GaloisField) AddOrSub(a, b int) int {
	_ = "STUB: not implemented"

	// Multiply multiplys two numbers
	return 0
}

func (gf *GaloisField) Multiply(a, b int) int { _ = "STUB: not implemented"; return 0 }

// Divide divides two numbers
func (gf *GaloisField) Divide(a, b int) int { _ = "STUB: not implemented"; return 0 }

func (gf *GaloisField) Invers(num int) int { _ = "STUB: not implemented"; return 0 }
