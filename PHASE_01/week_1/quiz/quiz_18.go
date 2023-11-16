package quiz

import (
	"fmt"
	"log"
	"strconv"
)

func Quiz_18() {
	num, err := prompt("input number: ", scanword, q18Valfn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(processDiamond(num))
}

func processDiamond(row int) string {
	// var mid int
	if row%2 != 0 {
		// mid =
	}
	return ""
}

func q18Valfn(input string) (int, error) {
	if len(input) < 1 {
		return -1, fmt.Errorf("you must provide value")
	}

	num, err := strconv.Atoi(input)
	if err != nil {
		return -1, fmt.Errorf("input must be number: %v", input)
	}

	if num < 1 {
		return -1, fmt.Errorf("input at least > 1")
	}

	return num, nil
}
