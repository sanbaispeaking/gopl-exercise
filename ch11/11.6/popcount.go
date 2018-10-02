package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// LookupCount returns the population count (number of set bits) of x.
func LookupCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// ShiftCount returns the population count (number of set bits) of x by shift x towards LSB 1 bit at a time.
func ShiftCount(x uint64) int {
	var count int
	for ; x > 0; x = x >> 1 {
		count += 1 & int(x)
	}
	return count
}

func ClearCount(x uint64) int {
	var count int
	for ; x > 0; x = x & (x - 1) {
		count++
	}
	return count
}
