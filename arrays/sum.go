package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(slices ...[]int) []int {
	// numOfSlices := len(slices)
	// println(numOfSlices)
	// sums := make([]int, numOfSlices)

	// for i, numbersSlice := range slices {
	// 	sums[i] = Sum(numbersSlice)
	// }
	// return sums
	var sums []int
	for _, numbers := range slices {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
