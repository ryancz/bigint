package bigint

import (
	"math/big"
	"strings"
)

func New(value int64) *big.Int {
	return big.NewInt(value)
}

func NewFromString(value string, base int) (*big.Int, bool) {
	if base == 16 {
		value = strings.TrimPrefix(strings.ToLower(value), "0x")
	}
	return New(0).SetString(value, base)
}

func Add(x, y *big.Int) *big.Int {
	return New(0).Add(x, y)
}

func Sub(x, y *big.Int) *big.Int {
	return New(0).Sub(x, y)
}

func Mul(x, y *big.Int) *big.Int {
	return New(0).Mul(x, y)
}

func Div(x, y *big.Int) *big.Int {
	return New(0).Div(x, y)
}