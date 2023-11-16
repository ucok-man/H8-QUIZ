package quiz

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Quiz_16() {
	performa, err := prompt("Input performa: ", scanword, q16performaValFn)
	if err != nil {
		log.Fatal(err)
	}

	gaji, err := prompt("Input gaji: ", scanword, q16gajiValFn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Bonus Karyawan: Rp.%.2f\n", calculateBonus(performa, gaji))
}

func calculateBonus(performa string, gaji float64) float64 {
	switch performa {
	case "A":
		return gaji * 20 / 100
	case "B":
		return gaji * 10 / 100
	case "C":
		return gaji * 5 / 100
	case "D":
		fallthrough
	default:
		return 0
	}
}

func q16performaValFn(input string) (string, error) {
	if len(input) < 1 {
		return "", fmt.Errorf("you must provide value")
	}

	validPerforma := []string{"A", "B", "C", "D"}
	inputCap := strings.ToUpper(input)

	var inputIsValid bool
	for _, pvalid := range validPerforma {
		if inputCap == pvalid {
			inputIsValid = true
		}
	}

	if !inputIsValid {
		return "", fmt.Errorf("performa must be in range [A, B, C, D]")
	}
	return inputCap, nil
}

func q16gajiValFn(input string) (float64, error) {
	if len(input) < 1 {
		return -1, fmt.Errorf("you must provide value")
	}

	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return -1, fmt.Errorf("you must provive valid number")
	}

	if num < 0 {
		return -1, fmt.Errorf("gaji must be positive")
	}
	return num, nil
}
