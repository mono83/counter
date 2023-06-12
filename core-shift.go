package counter

// Shift shifts counter values on one position right
func Shift[T Number](counter []T) {
	shift(0, counter)
}

// ShiftN shifts counter values to N positions right
func ShiftN[T Number](counter []T, n int) {
	// TODO make fast shifting
	if n > 32 {
		if n >= Capacity(counter) {
			initCounter(counter)
			return
		}
	}
	for i := 0; i < n; i++ {
		Shift(counter)
	}
}

func shift[T Number](offset int, counter []T) {
	if counter[offset+1] < 0 {
		// No value staged
		counter[offset+1] = counter[offset]
	} else {
		// Value already staged
		counter[offset+1] += counter[offset]
		if len(counter) > offset+2 {
			shift(offset+2, counter)
			counter[offset+2] = counter[offset+1]
		}
		counter[offset+1] = -1
	}
	counter[offset] = 0
}
