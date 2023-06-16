package counter

import "sync"

// NewCounter constructs new counter instance
func NewCounter[T Number](segments int) *Counter[T] {
	return &Counter[T]{values: Slice[T](segments)}
}

// Counter is concurrent-safe structure
type Counter[T Number] struct {
	m      sync.RWMutex
	values []T
}

func (c *Counter[T]) unsafeAdd(value T) {
	Add(c.values, value)
}

func (c *Counter[T]) unsafeShift() {
	Shift(c.values)
}

func (c *Counter[T]) unsafeShiftN(n int) {
	ShiftN(c.values, n)
}

func (c *Counter[T]) Add(value T) {
	c.m.Lock()
	c.unsafeAdd(value)
	c.m.Unlock()
}

func (c *Counter[T]) Inc() {
	c.Add(1)
}

func (c *Counter[T]) Shift() {
	c.m.Lock()
	c.unsafeShift()
	c.m.Unlock()
}

func (c *Counter[T]) ShiftN(n int) {
	c.m.Lock()
	c.unsafeShiftN(n)
	c.m.Unlock()
}

func (c *Counter[T]) Sum() (out T) {
	c.m.RLock()
	out = Sum(c.values)
	c.m.RUnlock()
	return
}

func (c *Counter[T]) SumN(segment int) (out T) {
	c.m.RLock()
	out = SumN(c.values, segment)
	c.m.RUnlock()
	return
}
