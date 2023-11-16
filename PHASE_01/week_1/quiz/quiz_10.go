package quiz

import (
	"fmt"
	"strconv"
)

func Quiz_10() {
	rows := 5

	for i := 0; i < rows; i++ {
		padding := strconv.Itoa(int(rows - i))
		fmt.Printf("%"+padding+"s", "")

		var num = 1
		for j := 0; j <= i; j++ {
			fmt.Printf("%2d", num)
			num = num * (i - j) / (j + 1)
		}
		fmt.Println()
	}
}
