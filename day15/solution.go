package day15

import "fmt"

// Run the solution.
func Run(inputPath string) {
	// numbers := []int{0, 3, 6}
	numbers := []int{8, 0, 17, 4, 1, 12}
	fmt.Printf("Part 1: %d\n", playGame(numbers, 2020))      // 981
	fmt.Printf("Part 2: %d\n", playGame2(numbers, 30000000)) // 164878
}

func playGame(numbers []int, iterations int) int {
	startSize := len(numbers)
	for i := 0; i < iterations-startSize; i++ {
		last := ageLast(numbers)
		numbers = append(numbers, last)
	}
	return numbers[len(numbers)-1]
}

func ageLast(numbers []int) int {
	last := numbers[len(numbers)-1]
	for i := len(numbers) - 2; i >= 0; i-- {
		if numbers[i] == last {
			return len(numbers) - 1 - i
		}
	}
	return 0
}

// part2 has 30m iterations so growing a slice to 60MB and searching gets slower
// the more numbers are added (even < 1m), so I switched to a map of each number
// to the last turn it was called in.
func playGame2(numbers []int, iterations int) int {
	last := 0
	prev := 0
	prevFound := false
	callByNumber := make(map[int]int)
	for turn := 0; turn < iterations; turn++ {
		if turn < len(numbers) {
			last = numbers[turn]
		} else {
			if prevFound {
				last = turn - prev - 1
			} else {
				last = 0
			}
		}
		prev, prevFound = callByNumber[last]
		callByNumber[last] = turn
	}
	return last
}
