package counter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	assert.Nil(t, Slice[int](0))
	assert.Nil(t, Slice[int](-1))
	if c := Slice[int](1); assert.NotNil(t, c) {
		assert.Equal(t, []int{0, -1}, c)
	}
	if c := Slice[int](5); assert.NotNil(t, c) {
		assert.Equal(t, []int{0, -1, 0, -1, 0, -1, 0, -1, 0, -1}, c)
	}
	if c := Slice[int](5); assert.NotNil(t, c) {
		assert.Equal(t, []int{0, -1, 0, -1, 0, -1, 0, -1, 0, -1}, c)
	}
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

func TestAt(t *testing.T) {
	// Empty slice
	assert.Equal(t, 0, At[int](nil, 0))
	assert.Equal(t, 0, At[int](nil, 1))

	// One segment
	assert.Equal(t, 8, At[int]([]int{8, -1}, 0))
	assert.Equal(t, 0, At[int]([]int{8, -1}, 1))

	// Two segments
	assert.Equal(t, 4, At[int]([]int{4, -1, 5, -1}, 0))
	assert.Equal(t, 5, At[int]([]int{4, -1, 5, -1}, 1))
	assert.Equal(t, 0, At[int]([]int{4, -1, 5, -1}, 2))

	// Two segments with staging
	assert.Equal(t, 4, At[int]([]int{4, 9, 5, 8}, 0))
	assert.Equal(t, 5, At[int]([]int{4, 9, 5, 8}, 1))
	assert.Equal(t, 0, At[int]([]int{4, 9, 5, 8}, 2))
}

func TestSum(t *testing.T) {
	assert.Equal(t, 0, Sum[int](nil))

	assert.Equal(t, 3, Sum[int]([]int{3, -1}))
	assert.Equal(t, 6, Sum[int]([]int{3, 3}))

	assert.Equal(t, 12, Sum[int]([]int{5, -1, 7, -1}))
	assert.Equal(t, 16, Sum[int]([]int{5, 3, 7, 1}))
}

func TestSumN(t *testing.T) {
	assert.Equal(t, 0, SumN[int](nil, 0))

	assert.Equal(t, 3, SumN[int]([]int{3, -1}, 0))
	assert.Equal(t, 3, SumN[int]([]int{3, 3}, 0))
	assert.Equal(t, 3, SumN[int]([]int{3, -1}, 1))
	assert.Equal(t, 3, SumN[int]([]int{3, 3}, 1))

	assert.Equal(t, 5, SumN[int]([]int{5, -1, 7, -1}, 0))
	assert.Equal(t, 12, SumN[int]([]int{5, -1, 7, -1}, 1))
	assert.Equal(t, 5, SumN[int]([]int{5, 3, 7, 1}, 0))
	assert.Equal(t, 15, SumN[int]([]int{5, 3, 7, 1}, 1))
}

func TestShift(t *testing.T) {
	c := []int{3, -1}
	Shift(c)
	assert.Equal(t, []int{0, 3}, c)

	c = []int{3, 0}
	Shift(c)
	assert.Equal(t, []int{0, -1}, c)

	c = []int{3, -1, 0, -1}
	Shift(c)
	assert.Equal(t, []int{0, 3, 0, -1}, c)

	c = []int{3, 2, 0, -1}
	Shift(c)
	assert.Equal(t, []int{0, -1, 5, 0}, c)
	Shift(c)
	assert.Equal(t, []int{0, 0, 5, 0}, c)
	Shift(c)
	assert.Equal(t, []int{0, -1, 0, -1}, c)
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

func TestIndexOf(t *testing.T) {
	assert.Equal(t, 0., IndexOf(0))
	assert.Equal(t, 1., IndexOf(1))
	assert.Equal(t, 1.5849625007211563, IndexOf(2))
	assert.Equal(t, 2., IndexOf(3))
	assert.Equal(t, 2.321928094887362, IndexOf(4))
	assert.Equal(t, 2.584962500721156, IndexOf(5))
	assert.Equal(t, 2.807354922057604, IndexOf(6))
	assert.Equal(t, 3., IndexOf(7))

	assert.Equal(t, 0, CeilIndexOf(0))
	assert.Equal(t, 1, CeilIndexOf(1))
	assert.Equal(t, 2, CeilIndexOf(2))
	assert.Equal(t, 2, CeilIndexOf(3))
	assert.Equal(t, 3, CeilIndexOf(4))
	assert.Equal(t, 3, CeilIndexOf(5))
	assert.Equal(t, 3, CeilIndexOf(6))
	assert.Equal(t, 3, CeilIndexOf(7))

	assert.Equal(t, 0, FloorIndexOf(0))
	assert.Equal(t, 1, FloorIndexOf(1))
	assert.Equal(t, 1, FloorIndexOf(2))
	assert.Equal(t, 2, FloorIndexOf(3))
	assert.Equal(t, 2, FloorIndexOf(4))
	assert.Equal(t, 2, FloorIndexOf(5))
	assert.Equal(t, 2, FloorIndexOf(6))
	assert.Equal(t, 3, FloorIndexOf(7))
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
