package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test16s(t *testing.T) {
	c := New[int](4)

	Add(c, 1)
	assert.Equal(t, 1, Sum(c))
	assert.Equal(t, []int{1, -1, 0, -1, 0, -1, 0, -1}, c)

	Shift(c)
	Add(c, 2)
	assert.Equal(t, 3, Sum(c))
	assert.Equal(t, []int{2, 1, 0, -1, 0, -1, 0, -1}, c)

	Shift(c)
	Add(c, 3)
	assert.Equal(t, 6, Sum(c))
	assert.Equal(t, []int{3, -1, 3, 0, 0, -1, 0, -1}, c)

	Shift(c)
	Add(c, 4)
	assert.Equal(t, 10, Sum(c))
	assert.Equal(t, []int{4, 3, 3, 0, 0, -1, 0, -1}, c)

	Shift(c)
	Add(c, 5)
	assert.Equal(t, 15, Sum(c))
	assert.Equal(t, []int{5, -1, 7, -1, 3, 0, 0, -1}, c)

	Shift(c)
	Add(c, 6)
	assert.Equal(t, 21, Sum(c))
	assert.Equal(t, []int{6, 5, 7, -1, 3, 0, 0, -1}, c)

	Shift(c)
	Add(c, 7)
	assert.Equal(t, 28, Sum(c))
	assert.Equal(t, []int{7, -1, 11, 7, 3, 0, 0, -1}, c)

	Shift(c)
	Add(c, 8)
	assert.Equal(t, 36, Sum(c))
	assert.Equal(t, []int{8, 7, 11, 7, 3, 0, 0, -1}, c)

	Shift(c)
	Add(c, 9)
	assert.Equal(t, 45, Sum(c))
	assert.Equal(t, []int{9, -1, 15, -1, 18, -1, 3, 0}, c)

	Shift(c)
	Add(c, 10)
	assert.Equal(t, 55, Sum(c))
	assert.Equal(t, []int{10, 9, 15, -1, 18, -1, 3, 0}, c)

	Shift(c)
	Add(c, 11)
	assert.Equal(t, 66, Sum(c))
	assert.Equal(t, []int{11, -1, 19, 15, 18, -1, 3, 0}, c)

	Shift(c)
	Add(c, 12)
	assert.Equal(t, 78, Sum(c))
	assert.Equal(t, []int{12, 11, 19, 15, 18, -1, 3, 0}, c)

	Shift(c)
	Add(c, 13)
	assert.Equal(t, 91, Sum(c))
	assert.Equal(t, []int{13, -1, 23, -1, 34, 18, 3, 0}, c)

	Shift(c)
	Add(c, 14)
	assert.Equal(t, 105, Sum(c))
	assert.Equal(t, []int{14, 13, 23, -1, 34, 18, 3, 0}, c)

	Shift(c)
	Add(c, 15)
	assert.Equal(t, 120, Sum(c))
	assert.Equal(t, []int{15, -1, 27, 23, 34, 18, 3, 0}, c)

	Shift(c)
	Add(c, 16)
	assert.Equal(t, 136, Sum(c))
	assert.Equal(t, []int{16, 15, 27, 23, 34, 18, 3, 0}, c)

	Shift(c)
	Add(c, 17)
	assert.Equal(t, 153-3, Sum(c))
	assert.Equal(t, []int{17, -1, 31, -1, 50, -1, 52, -1}, c)
}
