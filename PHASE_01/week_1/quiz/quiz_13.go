package quiz

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Quiz_13() {
	listnum, err := prompt("input numbers: ", scanline, q13numbersValfn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total: %v\n", sumNumbers(listnum...))
}

func sumNumbers(numbers ...int) int {
	var sum int
	for _, val := range numbers {
		sum += val
	}
	return sum
}

func q13numbersValfn(input string) ([]int, error) {
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
