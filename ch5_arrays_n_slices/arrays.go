package arrays

func SumCollection(collection []int) int {

	var sum int

	// range: iterate over collection. Returns: index, value
	// _ : ignore

	for _, num := range collection {
		sum += num
	}
	// equivalent:
	/*
		for i := range arr {
			sum += collection[i]
		}
	*/

	return sum
}

func untested() string {
	return "u didnt test me"
}

// SumAll returns a slice with the sums of received slices (variadic)
func SumAll(slices ...[]int) []int {

	var sums []int

	for _, slice := range slices {
		sums = append(sums, SumCollection(slice))
	}

	return sums
}

// SumAllTails returns a slice with the sums of the tails of given slices (variadic)
func SumAllTails(slices ...[]int) []int {

	var sums []int

	for _, slice := range slices {
		if len(slice) > 1 {
			// slice[low:high]
			tails := slice[1:]

			sums = append(sums, SumCollection(tails))
		} else if len(slice) > 0 {
			sums = append(sums, slice[0])
		} else {
			sums = append(sums, 0)
		}
	}

	return sums
}
