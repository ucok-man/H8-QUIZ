package quiz

import (
	"fmt"
	"os"
	"strings"
)

func Quiz_09() {
	input, err := getInput("Masukan dimensi (min 1): ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	asterisk := procces(input)
	fmt.Println(asterisk)

}

func getInput(prompt string) (int, error) {
	fmt.Print(prompt)

	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		return -1, fmt.Errorf("cannot procces your input: make sure enter valid number")
	}

	if num < 1 {
		return -1, fmt.Errorf("cannot procces your input: number must be positive and min = 1")
	}

	return num, nil
}

func procces(input int) string {
	var result string
	for i := 1; i <= input; i++ {
		if i == 1 {
			result += strings.Repeat("*", input)
			result += "\n"
			continue
		}

		if i == input {
			result += strings.Repeat("*", input)
			result += "\n"
			continue
		}

		var format string
		format += "*"
		format += strings.Repeat(" ", input-2)
		format += "*"
		result += format + "\n"
	}
	return result
}
