package quiz

import "fmt"

func Quiz_03() {
	var number int

	if number%2 == 0 {
		fmt.Println("number adalah ganjil")
		return
	}
	fmt.Println("number adalah genap")
}
