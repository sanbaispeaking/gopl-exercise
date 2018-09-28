package popcount

// PopCount returns the population count (number of set bits) of x by shift x towards LSB 1 bit at a time.
func PopCount(x uint64) int {
	var count int
	for i := 0; i < 64; i++ {
		count += int((x >> uint(i)) & 1)
	}
	return count
}
