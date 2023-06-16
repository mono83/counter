package counter

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestDurationCounter(t *testing.T) {
	// 5 segments 10 seconds each is 310 seconds, slightly more than 5 minutes
	c := NewDurationCounter[int](time.Second*10, 5)
	assert.True(t, c.last.IsZero())
	assert.Equal(t, []int{0, -1, 0, -1, 0, -1, 0, -1, 0, -1}, c.counter.values)

	c.AddAt(time.Unix(18, 0), 3) // 18-th second
	assert.Equal(t, time.Unix(10, 0), c.last)
	assert.Equal(t, []int{3, -1, 0, -1, 0, -1, 0, -1, 0, -1}, c.counter.values)
	assert.Equal(t, 3, c.CeilSumD(time.Second*10))
	assert.Equal(t, 3, c.CeilSumD(time.Second*300))

	c.AddAt(time.Unix(18, 0), 2) // Same second
	assert.Equal(t, time.Unix(10, 0), c.last)
	assert.Equal(t, []int{5, -1, 0, -1, 0, -1, 0, -1, 0, -1}, c.counter.values)
	assert.Equal(t, 5, c.CeilSumD(time.Second*10))
	assert.Equal(t, 5, c.CeilSumD(time.Second*300))

	c.AddAt(time.Unix(24, 0), 8) // 24-th second
	assert.Equal(t, time.Unix(20, 0), c.last)
	assert.Equal(t, []int{8, 5, 0, -1, 0, -1, 0, -1, 0, -1}, c.counter.values)
	assert.Equal(t, 8, c.CeilSumD(time.Second*10))
	assert.Equal(t, 13, c.CeilSumD(time.Second*300))

	c.AddAt(time.Unix(36, 0), 2) // 36-th second
	assert.Equal(t, time.Unix(30, 0), c.last)
	assert.Equal(t, []int{2, -1, 13, 0, 0, -1, 0, -1, 0, -1}, c.counter.values)
	assert.Equal(t, 2, c.CeilSumD(time.Second*10))
	assert.Equal(t, 15, c.CeilSumD(time.Second*14))
	assert.Equal(t, 15, c.CeilSumD(time.Second*300))
}

func TestDurationCounterConcurrency(t *testing.T) {
	n := 1000
	wg := sync.WaitGroup{}
	wg.Add(n)

	at := time.Now()
	c := NewDurationCounter[int](time.Second, 5)
	for i := 0; i < n; i++ {
		go func() {
			c.AddAt(at, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	assert.Equal(t, n, c.Sum())
}
