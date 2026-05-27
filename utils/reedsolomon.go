package utils

import (
	"sync"
)

type ReedSolomonEncoder struct {
	gf        *GaloisField
	polynomes []*GFPoly
	m         *sync.Mutex
}

func NewReedSolomonEncoder(gf *GaloisField) *ReedSolomonEncoder {
	_ = "STUB: not implemented"
	return nil
}

func (rs *ReedSolomonEncoder) getPolynomial(degree int) *GFPoly {
	_ = "STUB: not implemented"
	return nil
}

func (rs *ReedSolomonEncoder) Encode(data []int, eccCount int) []int {
	_ = "STUB: not implemented"
	return nil
}
