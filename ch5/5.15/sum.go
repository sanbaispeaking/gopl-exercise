package main

// Min returns the minimum item of a sequence of integers
func Min(items ...int) int {
	if len(items) == 0 {
		panic("no values to compare")
	}

	min := items[0]
	for _, item := range items {
		if min > item {
			min = item
		}
	}
	return min
}

// Max returns the maximum item of a sequence of intergers
func Max(items ...int) int {
	if len(items) == 0 {
		panic("no values to compare")
	}

	max := items[0]
	for _, item := range items {
		if max < item {
			max = item
		}
	}
	return max
}

func main() {

}
