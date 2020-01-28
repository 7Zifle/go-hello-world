package arrays

func Sum(numbers []int) int {
	var total int
	for _, value := range numbers {
		total += value
	}
	return total
}

func SumAll(numbers ...[]int) []int {
	var sums []int

	for _, value := range numbers {
		sums = append(sums, Sum(value))
	}

	return sums
}

func SumAllTails(numbers ...[]int) []int {
	var tails []int

	for _, value := range numbers {
		if len(value) == 0 {
			tails = append(tails, 0)
		} else {
			tail := value[1:]
			tails = append(tails, Sum(tail))
		}
	}
	return tails
}
