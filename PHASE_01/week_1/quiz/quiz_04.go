package quiz

import (
	"fmt"
	"os"
)

func Quiz_04() {
	fmt.Print("Masukan Bulan (number): ")

	var indexBulan int
	_, err := fmt.Scan(&indexBulan)
	if err != nil {
		fmt.Fprintln(os.Stderr, "cannot procces your input: make sure enter valid number")
		os.Exit(1)
	}

	if indexBulan < 0 || indexBulan > 12 {
		fmt.Fprintln(os.Stderr, "cannot procces your input: input must be 1 - 12")
		os.Exit(1)
	}

	bulans := []string{"January", "Febuary", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}

	fmt.Printf("bulan ke %d: %s\n", indexBulan, bulans[indexBulan-1])

}
