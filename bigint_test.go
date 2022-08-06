package bigint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	zero := New(0)
	assert.Equal(t, int64(0), zero.Int64())
	one := New(1)
	assert.Equal(t, int64(1), one.Int64())
}

func TestNewFromString(t *testing.T) {
	bi, ok := NewFromString("1", 10)
	assert.True(t, ok)
	assert.Equal(t, int64(1), bi.Int64())

	s := "1000000000000000000000000000000"
	bi, ok = NewFromString(s, 10)
	assert.True(t, ok)
	assert.Equal(t, s, bi.String())

	bi, ok = NewFromString("0x10", 16)
	assert.True(t, ok)
	assert.Equal(t, int64(16), bi.Int64())

	s = "0x1234567890abcdef1234567890"
	bi, ok = NewFromString(s, 16)
	assert.True(t, ok)
	assert.Equal(t, s[2:], bi.Text(16))
}

func TestAdd(t *testing.T) {
	one := New(1)
	two := New(2)
	three := Add(one, two)
	assert.Equal(t, int64(3), three.Int64())

	s := "1000000000000000000000000000000"
	bi1, ok := NewFromString(s, 10)
	assert.True(t, ok)
	bi2 := Add(bi1, one)
	assert.Equal(t, "1000000000000000000000000000001", bi2.String())
}

func TestSub(t *testing.T) {
	one := New(1)
	two := New(2)
	negOne := Sub(one, two)
	assert.Equal(t, int64(-1), negOne.Int64())

	s := "1000000000000000000000000000001"
	bi1, ok := NewFromString(s, 10)
	assert.True(t, ok)
	bi2 := Sub(bi1, one)
	assert.Equal(t, "1000000000000000000000000000000", bi2.String())
}

func TestMul(t *testing.T) {
	zero := New(0)
	ten := New(10)
	bi1 := Mul(zero, ten)
	assert.Equal(t, int64(0), bi1.Int64())

	bi1 = New(1e10)
	bi2 := New(1e11)
	bi3 := Mul(bi1, bi2)
	assert.Equal(t, "1000000000000000000000", bi3.String())
}

func TestDiv(t *testing.T) {
	one := New(1)
	ten := New(10)
	bi := Div(ten, one)
	assert.Equal(t, int64(10), bi.Int64())

	s := "1000000000000000000000"
	bi1, ok := NewFromString(s, 10)
	assert.True(t, ok)
	bi2 := New(1e10)
	bi3 := Div(bi1, bi2)
	assert.Equal(t, int64(1e11), bi3.Int64())
}