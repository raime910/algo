package main

import "fmt"

func main() {
	result := twoNumberSum([]int{3, 5, -4, 8, 11, 1, -1, 6}, 10)

	fmt.Println(result)
}

func twoNumberSum(array []int, target int) []int {
	// Write your code here.
	numMap := map[int]bool{}

	for _, num := range array {
		difference := target - num

		_, ok := numMap[num]

		fmt.Println(numMap)
		fmt.Println(ok)

		if ok {
			return []int{difference, num}
		}

		numMap[difference] = true
	}

	return nil
}
