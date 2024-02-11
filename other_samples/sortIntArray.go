package main

import (
	"fmt"
	"sort"
)

func SortNumbers(numbers []int) []int {
	if len(numbers) == 0 {
		return []int{}
	} else {
		sort.Ints(numbers)
		return numbers
	}
}

func main() {
	nums := []int{4, 6, 3, 1, 5, 8}
	s := SortNumbers(nums)
	fmt.Println(s)
}
