package quiz

import (
	"bufio"
	"fmt"
	"os"
)

type scanmode int

const (
	scanword scanmode = 1
	scanline scanmode = 2
)

func prompt[T any](perintah string, mode scanmode, validation func(input string) (T, error)) (T, error) {
	var result T

	fmt.Print(perintah)

	scanner := bufio.NewScanner(os.Stdin)
	if mode == scanword {
		scanner.Split(bufio.ScanWords)
	}
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return result, fmt.Errorf("error when scanning input: %v", err)
	}
	return validation(scanner.Text())
}
