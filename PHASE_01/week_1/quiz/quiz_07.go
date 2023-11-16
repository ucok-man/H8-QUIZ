package quiz

import "fmt"

func Quiz_07() {
	for i := 1; i <= 20; i++ {
		if i%3 == 0 {
			fmt.Print("fizz ")
			continue
		}

		if i%5 == 0 {
			fmt.Print("buzz ")
			continue
		}

		fmt.Printf("%d ", i)
	}
	fmt.Println()
}
