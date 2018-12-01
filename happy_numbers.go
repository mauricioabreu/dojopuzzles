package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	number, _ := strconv.Atoi(os.Args[1])
	fmt.Println(IsHappyNumber(number))
}

// IsHappyNumber : find out if the given number is a happy number
// Given the number 7:
// 7² = 49
// 4² + 9² = 97
// 9² + 7² = 130
// 1² + 3² + 0² = 10
// 1² + 0² = 1
func IsHappyNumber(number int) bool {
	return isHappyNumber(number, []int{})
}

func isHappyNumber(number int, seen []int) bool {
	accum := 0
	for number > 0 {
		remaining := number % 10
		number = number / 10
		accum += remaining * remaining
	}
	// If the number was already seen it means
	// we entered an infinite looping, resulting
	// in a sad number :(
	if wasSeen(accum, seen) {
		return false
	}
	seen = append(seen, accum)
	if accum == 1 {
		return true
	}
	return isHappyNumber(accum, seen)
}

func wasSeen(number int, numbers []int) bool {
	for _, value := range numbers {
		if number == value {
			return true
		}
	}
	return false
}
