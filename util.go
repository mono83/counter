package counter

func initialize[T Number](counter []T) {
	seg := segmentCount(counter)
	for i := 0; i < seg; i++ {
		counter[(i+1)*2-1] = -1
	}
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
