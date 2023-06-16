package counter

import (
	"fmt"
	"math"
)

// Slice creates new counter slice with given amount of slots
// Each slot contains two data cells and one staging cell
func Slice[T Number](segments int) []T {
	if segments < 1 {
		return nil
	}
	counter := make([]T, segments*2)
	initialize(counter)
	return counter
}

// Add adds value to counter
func Add[T Number](counter []T, value T) {
	counter[0] += value
}

// Inc increments value in counter
func Inc[T Number](counter []T) {
	Add(counter, 1)
}

// Sum calculates sum of all values in counter slice
func Sum[T Number](counter []T) (sum T) {
	l := len(counter)
	for i := 0; i < l; i++ {
		if i%2 == 0 {
			sum += counter[i]
		} else if counter[i] > 0 {
			sum += counter[i]
		}
	}
	return
}

// SumN calculates value of all N segments in counter slice
// starting with zero.
func SumN[T Number](counter []T, segment int) (sum T) {
	i := 0
	l := len(counter)
	for segment >= 0 {
		if i == l {
			break
		}
		sum += counter[i]
		if counter[i+1] > 0 && i+1 != l-1 && segment != 0 { // Enabled, not negative, not last
			sum += counter[i+1]
		}
		segment--
		i += 2
	}
	return
}

// At returns counter value at given segment position
func At[T Number](counter []T, segment int) (value T) {
	j := segment * 2
	if l := len(counter); l > j {
		value = counter[j]
	}
	return
}

// IndexOf returns matching segment index as float
func IndexOf(value int) float64 {
	return math.Log2(float64(value + 1))
}

// CeilIndexOf returns matching segment index rounding up
func CeilIndexOf(value int) int {
	switch value {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 2
	case 4:
		return 3
	default:
		return int(math.Ceil(IndexOf(value)))
	}
}

// FloorIndexOf returns matching segment index rounding down
func FloorIndexOf(value int) int {
	return int(math.Floor(IndexOf(value)))
}

// Capacity returns maximum capacity of counter slice
func Capacity[T Number](counter []T) (cap int) {
	for i := 0; i < segmentCount(counter); i++ {
		cap += segmentMultiplier(i)
	}
	return
}

// Shift shifts counter values on one position right
func Shift[T Number](counter []T) {
	shift(0, counter)
}

// ShiftN shifts counter values to N positions right
func ShiftN[T Number](counter []T, n int) {
	// TODO make fast shifting
	if n > 32 {
		if n >= Capacity(counter) {
			initialize(counter)
			return
		}
	}
	for i := 0; i < n; i++ {
		Shift(counter)
	}
}

// Print outputs counter slice contents to standard output
func Print[T Number](counter []T) {
	segments := segmentCount(counter)
	fmt.Printf("Counter with %d segments, capacity %d\n", segments, Capacity(counter))
	for i := 0; i < segments; i++ {
		j := i * 2
		fmt.Printf("%02dx: %d\n", int(math.Pow(2, float64(i))), counter[j])
		if counter[j+1] < 0 {
			fmt.Println("---: not staged")
		} else {
			fmt.Println("---:", counter[j+1])
		}
	}
}
