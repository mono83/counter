package counter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	assert.Nil(t, Slice[int](0))
	assert.Nil(t, Slice[int](-1))
	if c := Slice[int](5); assert.NotNil(t, c) {
		assert.Len(t, c, 10)
	}

	assert.Equal(t, 0, Sum(Slice[int](0)))
	assert.Equal(t, 0, Sum(Slice[int](-1)))
	assert.Equal(t, 0, Sum(Slice[int](5)))
}

func TestAddInc(t *testing.T) {
	c := Slice[int](1)
	Add(c, 5)
	assert.Equal(t, 5, c[0])
	Add(c, 8)
	assert.Equal(t, 13, c[0])
	assert.Equal(t, 13, c[0])
	Inc(c)
	assert.Equal(t, 14, c[0])
	Inc(c)
	assert.Equal(t, 15, c[0])

	assert.Equal(t, 15, Sum(c))
}

func TestSumN(t *testing.T) {
	c := Slice[int](4)
	Add(c, 5)
	Shift(c)
	Add(c, 3)

	assert.Equal(t, 3, SumN(c, 0))
	assert.Equal(t, 8, SumN(c, 1))
	assert.Equal(t, 8, SumN(c, 2))   // Overflow
	assert.Equal(t, 8, SumN(c, 100)) // Overflow

	Shift(c)
	Add(c, 10)
	assert.Equal(t, 10, SumN(c, 0))
	assert.Equal(t, 18, SumN(c, 1))
	assert.Equal(t, 18, SumN(c, 2))
	assert.Equal(t, 18, SumN(c, 100)) // Overflow
}

func TestShift(t *testing.T) {
	c := Slice[int](3)
	Add(c, 2)
	assert.Equal(t, 2, c[0])
	Shift(c)
	assert.Equal(t, 0, c[0])
	Add(c, 8)
	assert.Equal(t, 8, c[0])
	Shift(c)
	assert.Equal(t, 10, c[2])
	assert.Equal(t, 10, Sum(c))
	Add(c, 3)
	assert.Equal(t, 3, c[0])
	Shift(c)
	Add(c, 9)
	Shift(c)
	assert.Equal(t, 22, Sum(c))
}

func TestShiftN(t *testing.T) {
	c1 := Slice[int](15)

	Inc(c1)
	for i := 0; i < 512; i++ {
		Shift(c1)
		c2 := Slice[int](15)
		Inc(c2)
		ShiftN(c2, i+1)
		assert.Equal(t, c1, c2)
	}
}

func TestCapacity(t *testing.T) {
	assert.Equal(t, 1, Capacity(Slice[int](1)))      // 1
	assert.Equal(t, 3, Capacity(Slice[int](2)))      // 1 2
	assert.Equal(t, 7, Capacity(Slice[int](3)))      // 1 2 4
	assert.Equal(t, 15, Capacity(Slice[int](4)))     // 1 2 4 8
	assert.Equal(t, 31, Capacity(Slice[int](5)))     // 1 2 4 8 16
	assert.Equal(t, 63, Capacity(Slice[int](6)))     // 1 2 4 8 16 32
	assert.Equal(t, 127, Capacity(Slice[int](7)))    // 1 2 4 8 16 32 64
	assert.Equal(t, 255, Capacity(Slice[int](8)))    // 1 2 4 8 16 32 64 128
	assert.Equal(t, 511, Capacity(Slice[int](9)))    // 1 2 4 8 16 32 64 128 256
	assert.Equal(t, 1023, Capacity(Slice[int](10)))  // 1 2 4 8 16 32 64 128 256 512
	assert.Equal(t, 2047, Capacity(Slice[int](11)))  // 1 2 4 8 16 32 64 128 256 512 1024
	assert.Equal(t, 4095, Capacity(Slice[int](12)))  // 1 2 4 8 16 32 64 128 256 512 1024 2048
	assert.Equal(t, 8191, Capacity(Slice[int](13)))  // 1 2 4 8 16 32 64 128 256 512 1024 2048 4096
	assert.Equal(t, 16383, Capacity(Slice[int](14))) // 1 2 4 8 16 32 64 128 256 512 1024 2048 4096 8192
	assert.Equal(t, 32767, Capacity(Slice[int](15))) // 1 2 4 8 16 32 64 128 256 512 1024 2048 4096 8192 16384
}
