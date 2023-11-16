package quiz

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Quiz_20() {
	// listNilai := []int{85, 60, 78, 92, 45, 73}

	listNilai, err := prompt("input nilai: ", scanline, valfn)
	if err != nil {
		log.Fatal(err)
	}
	hasil := evaluasi(listNilai)

	// print output
	for _, val := range hasil {
		fmt.Println(val)
	}

}

func evaluasi(listNilai []int) []string {
	result := []string{}
	for idx, nilai := range listNilai {
		switch {
		case nilai >= 85:
			result = append(result, fmt.Sprintf("Siswa %d mendapat predikat A", idx+1))
		case nilai >= 70:
			result = append(result, fmt.Sprintf("Siswa %d mendapat predikat B", idx+1))
		case nilai >= 50:
			result = append(result, fmt.Sprintf("Siswa %d mendapat predikat C", idx+1))
		case nilai < 50:
			result = append(result, fmt.Sprintf("Siswa %d mendapat predikat D", idx+1))
		default:
			panic(fmt.Sprintf("unhandled condition %v", nilai))
		}
	}
	return result
}

func valfn(input string) ([]int, error) {
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

		if n < 0 {
			return nil, fmt.Errorf("input must be positive: %v", num)
		}
		result = append(result, n)
	}
	return result, nil
}
