package quiz

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func Quiz_14() {
	listnum, err := prompt("input numbers: ", scanline, q14numbersValfn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Max num: %v\n", findmax(listnum...))
}

func findmax(numbers ...int) int {
	max := math.MinInt
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}

	return max
}

func q14numbersValfn(input string) ([]int, error) {
	inputs := strings.Split(input, " ")
	if len(inputs) < 1 {
		return nil, fmt.Errorf("you must provide value")
	}

	var result = []int{}
	for _, num := range inputs {
		n, err := strconv.Atoi(num)
		if err != nil {
			return nil, fmt.Errorf("input must be number: %v", num)
		}

		result = append(result, n)
	}
	return result, nil
}
