package counter

import "sync"

type Counter[T Number] struct {
	m sync.RWMutex
	values []T
}

func (c *Counter[T]) Add(value T) {
	c.m.Lock()
	Add(c.values, value)
	c.m.Unlock()
}

func (c *Counter[T]) Inc() {
	c.Add(1)
}

func (c *Counter[T]) Shift() {
	c.m.Lock()
	Shift(c.values)
	c.m.Unlock()
}

func (c *Counter[T]) ShiftN(n int) {
	c.m.Lock()
	ShiftN(c.values, n)
	c.m.Unlock()
}

func (c *Counter[T]) Sum() (sum T) {
	c.m.RLock()
	sum = Sum(c.values)
	c.m.RUnlock()
	return
}

func (c *Counter[T]) SumAt(index int) (sum T) {
	c.m.RLock()
	sum = SumN(c.values, index)
	c.m.RUnlock()
	return
}