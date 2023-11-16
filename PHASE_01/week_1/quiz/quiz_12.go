package quiz

import (
	"fmt"
	"log"
	"strconv"
)

func Quiz_12() {
	num1, err := prompt("input number 1: ", scanword, q12Valfn)
	if err != nil {
		log.Fatal(err)
	}

	num2, err := prompt("input number: ", scanword, q12Valfn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Sum: %v\n", sumnumber(num1, num2))
}

func sumnumber(a, b int) int {
	return a + b

}

func q12Valfn(input string) (int, error) {
	if len(input) < 1 {
		return -1, fmt.Errorf("you must provide value")
	}

	num, err := strconv.Atoi(input)
	if err != nil {
		return -1, fmt.Errorf("input must be number: %v", input)
	}
	return num, nil
}
