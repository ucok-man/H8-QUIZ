package quiz

import (
	"fmt"
	"log"
	"strings"
)

func joinStr(inputs ...string) (string, int) {
	var result string
	for _, str := range inputs {
		result += " " + str
	}
	return result, len(inputs)
}

func Quiz_19() {
	listword, err := prompt("inputs many words: ", scanline, q19valfn)
	if err != nil {
		log.Fatal(err)
	}

	result, n := joinStr(listword...)
	fmt.Printf("result: %v\n", result)
	fmt.Printf("num word joined: %v\n", n)
}

func q19valfn(input string) ([]string, error) {
	inputs := strings.Split(input, " ")
	if len(inputs) < 1 {
		return nil, fmt.Errorf("you must provide value")
	}
	return inputs, nil
}
