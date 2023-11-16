package quiz

import "fmt"

func sum(a, b int) int {
	return a + b
}

func Quiz_11() {
	var a, b = 10, 20
	fmt.Println("Sum number: ", sum(a, b))
}
