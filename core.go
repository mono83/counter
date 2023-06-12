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
	initCounter(counter)
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
	for i := 0; i < len(counter); i++ {
		if counter[i] < 0 {
			continue
		}
		sum += counter[i]
	}
	return
}

// SumN calculates value of all N values in counter slice
// starting with zero.
func SumN[T Number](counter []T, index int) (sum T) {
	i := 0
	for index >= 0 {
		if i == len(counter) {
			break
		}
		if counter[i] > 0 {
			sum += counter[i]
			index--
		}
		i++
	}
	return
}

func segmentCount[T Number](counter []T) int {
	return len(counter) / 2
}

func segmentMultiplier(segment int) (m int) {
	m = 1
	for i := 0; i < segment; i++ {
		m *= 2
	}
	return
}

func initCounter[T Number](counter []T) {
	seg := segmentCount(counter)
	for i := 0; i < seg; i++ {
		counter[(i+1)*2-1] = -1
	}
}

func Capacity[T Number](counter []T) (cap int) {
	for i := 0; i < segmentCount(counter); i++ {
		cap += segmentMultiplier(i)
	}
	return
}

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
