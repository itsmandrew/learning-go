package arrays_slices

func Sum(nums [5]int) int {
	sum := 0

	for i := 0; i < 5; i++ {
		sum += nums[i]
	}
	return sum
}

// Keep in mind if I wanted to pass [4]int intto a function that
// expects [5]int it won't compile. Different types

// Also keep in mind most of the time we use slices in Go.

// Refactored version, using range
func Sum_v2(numbers [5]int) int {
	sum := 0

	// range returns two values, index and value
	for _, num := range numbers {
		sum += num
	}

	return sum
}

// Function that accepts one SLICE
func SumDynamic(numbers []int) int {
	sum := 0

	for _, value := range numbers {
		sum += value
	}
	return sum
}

// Go does not let u use equality operators with slices.
// We could write a function to iterate over each got and want slice and check their values
// We aren't going to do that and just use 'reflect.DeepEqual
func SumAll(numbers ...[]int) []int {

	lengthOfNumbers := len(numbers)

	// Another way to create a slice, starting cap of len
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbers {
		sums[i] = SumDynamic(numbers)
	}

	return sums
}

// In this implementation, we worry less about. We start w/ an empty slice and append
func SumAll_v2(numbers ...[]int) []int {
	var sums []int

	for _, nums := range numbers {
		sums = append(sums, SumDynamic(nums))
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
			sums = append(sums, SumDynamic(tail))
		}
	}

	return sums
}
