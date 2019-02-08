package main

import "fmt"

func main() {
	input := []int{100, 101, 102, 103, 104, 105, 110, 111, 113, 114, 115, 150}
	fmt.Println(GetIntervals(input))
}

// GetIntervals return a list of numbers grouped in intervals
// Given the input array [1, 2, 5, 6, 7, 10]
// Return a new array ["1-2", "5-7", "10"]
func GetIntervals(numbers []int) []string {
	ranges := make([][]int, 0)

	if len(numbers) == 0 {
		return []string{}
	}

	// Iterate over numbers, grouping them in the array
	previous := numbers[0]
	for _, number := range numbers {
		if number == previous+1 {
			ranges[len(ranges)-1] = append(ranges[len(ranges)-1], number)
		} else {
			ranges = append(ranges, []int{number})
		}
		previous = number
	}

	return formatIntervals(ranges)
}

func formatIntervals(ranges [][]int) []string {
	intervals := make([]string, 0)

	for _, interval := range ranges {
		if len(interval) > 1 {
			intervals = append(intervals, fmt.Sprintf("%d-%d", interval[0], interval[len(interval)-1]))
		} else {
			intervals = append(intervals, fmt.Sprintf("%d", interval[0]))
		}
	}
	return intervals
}
