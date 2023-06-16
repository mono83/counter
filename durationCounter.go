package counter

import (
	"time"
)

func NewDurationCounter[T Number](slot time.Duration, segments int) *DurationCounter[T] {
	return &DurationCounter[T]{
		counter: NewCounter[T](segments),
		slot:    slot,
	}
}

type DurationCounter[T Number] struct {
	counter *Counter[T]
	slot    time.Duration
	last    time.Time
}

func (d *DurationCounter[T]) AddAt(t time.Time, value T) {
	t = t.Truncate(d.slot)
	d.counter.m.Lock()
	if d.last != t {
		delta := int(t.Sub(d.last) / d.slot)
		if delta > 0 && !d.last.IsZero() {
			d.counter.unsafeShiftN(delta)
		}
		d.last = t
	}
	d.counter.unsafeAdd(value)
	d.counter.m.Unlock()
}

func (d *DurationCounter[T]) Add(value T) {
	d.AddAt(time.Now(), value)
}

func (d *DurationCounter[T]) Inc() {
	d.AddAt(time.Now(), 1)
}

func (d *DurationCounter[T]) Sum() T {
	return d.counter.Sum()
}

func (d *DurationCounter[T]) CeilSumD(delta time.Duration) T {
	segment := 0
	if delta > d.slot {
		segment = CeilIndexOf(int(delta / d.slot))
	}
	return d.counter.SumN(segment)
}

func (d *DurationCounter[T]) FloorSumD(delta time.Duration) T {
	segment := 0
	if delta > d.slot {
		segment = FloorIndexOf(int(delta / d.slot))
	}
	return d.counter.SumN(segment)
}
