package main

// Sum takes an array of integers and returns their sum
func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}

	return
}

// SumAll takes a slice of integer slices and returns a slice containing the sums of each integer slice
func SumAll(collections ...[]int) (sums []int) {
	for _, collection := range collections {
		sums = append(sums, Sum(collection))
	}

	return
}

// SumAllTails takes a slice of integer slices and returns a slice containing the sums of the tails of each integer slice
func SumAllTails(collections ...[]int) (sums []int) {
	for _, collection := range collections {
		if len(collection) <= 1 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(collection[1:]))
		}
	}

	return
}

func main() {}
