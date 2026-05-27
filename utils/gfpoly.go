package utils

type GFPoly struct {
	gf           *GaloisField
	Coefficients []int
}

func (gp *GFPoly) Degree() int { _ = "STUB: not implemented"; return 0 }

func (gp *GFPoly) Zero() bool { _ = "STUB: not implemented"; return false }

// GetCoefficient returns the coefficient of x ^ degree
func (gp *GFPoly) GetCoefficient(degree int) int { _ = "STUB: not implemented"; return 0 }

func (gp *GFPoly) AddOrSubstract(other *GFPoly) *GFPoly { _ = "STUB: not implemented"; return nil }

func (gp *GFPoly) MultByMonominal(degree int, coeff int) *GFPoly {
	_ = "STUB: not implemented"
	return nil
}

func (gp *GFPoly) Multiply(other *GFPoly) *GFPoly { _ = "STUB: not implemented"; return nil }

func (gp *GFPoly) Divide(other *GFPoly) (quotient *GFPoly, remainder *GFPoly) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewMonominalPoly(field *GaloisField, degree int, coeff int) *GFPoly {
	_ = "STUB: not implemented"
	return nil
}

func NewGFPoly(field *GaloisField, coefficients []int) *GFPoly {
	_ = "STUB: not implemented"
	return nil
}
