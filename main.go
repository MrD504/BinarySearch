package main


import (
	"math"
	"fmt"
)

var (
	targetValue int = 73
	primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37,
		41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
)

func main() {
	result := doSearch(primes, targetValue)
	fmt.Printf("The value %v was found at index %v \n", targetValue, result)
}

func doSearch(array []int, targetValue int) int64{
	min := 0.0
	max := float64(len(array) -1)
	
	for max >= min {
		var average float64 = (min + max)/ 2.0
		guessf64 := math.Floor(average)
		guess := int64(guessf64)
		if (array[guess] == targetValue) {
			return int64(guess)
		} else if (array[guess] < targetValue) {
			min = float64(guess) + 1.0
		} else {
			max = float64(guess) - 1.0
		}
	}

	return -1
}