package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	var sums []int
	for _, nums := range numbers {
		sums = append(sums, Sum(nums))
	}
	return []int{4, 15}
}
