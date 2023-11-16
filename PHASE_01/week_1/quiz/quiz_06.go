package quiz

import (
	"fmt"
	"os"
)

func Quiz_06() {
	fmt.Print("Index geometri: ")

	var idx int
	_, err := fmt.Scan(&idx)
	if err != nil {
		fmt.Fprintln(os.Stderr, "cannot procces your input: make sure enter valid number")
		os.Exit(1)
	}

	if idx <= 0 {
		fmt.Fprintln(os.Stderr, "cannot procces your input: index must be positive greater than 0")
		os.Exit(1)
	}

	var result = 1
	for i := 1; i <= idx; i++ {
		if i == 1 {
			fmt.Printf("%d ", result)
			continue
		}
		result *= 2
		fmt.Printf("%d ", result)
	}
	fmt.Println()
}
