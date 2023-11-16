package quiz

import "fmt"

type Book struct {
	judul   string
	penulis string
	harga   float64
}

func Quiz_15() {
	books := []Book{
		{
			judul:   "Nobita",
			penulis: "Nobura",
			harga:   20.00,
		},
		{
			judul:   "Nobita",
			penulis: "Nobura",
			harga:   20.00,
		}, {
			judul:   "Nobita",
			penulis: "Nobura",
			harga:   20.00,
		}, {
			judul:   "Nobita",
			penulis: "Nobura",
			harga:   20.00,
		}, {
			judul:   "Nobita",
			penulis: "Nobura",
			harga:   20.00,
		},
	}
	fmt.Printf("books: %v\n", books)
	fmt.Printf("Total harga: %.2f\n", totalHarga(books))
}

func totalHarga(books []Book) float64 {
	var sum float64
	for _, book := range books {
		sum += book.harga
	}
	return sum
}
