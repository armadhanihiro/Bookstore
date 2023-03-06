package main

import (
	"bookstore/model"
	"bookstore/service"
	"fmt"
)
func main() {
	new_book := []model.Book{
		{
			ID:         1,
			Titles:     "Laskar Pelangi",
			Authors:    "Andrea Hirata",
			Quantities: 3,
			Available:  true,
		},
		{
			ID:         2,
			Titles:     "Ranah 3 Warna",
			Authors:    "Ahmad Fuadi",
			Quantities: 1,
			Available:  true,
		},
	}

	bookController := service.BookStoreDB{}

	insertBook := bookController.InsertBook(&new_book, bookController)
	if insertBook != nil {
		fmt.Println("Can't Add Book")
		return
	}

	// b, err := bookController.GetBookByID(2, bookController)
	// if err != nil {
	// 	fmt.Println("Can't get book with ID = 2")
	// 	return
	// }
	// fmt.Println("Book with ID 2 : ", b)

	// err := bookController.DeleteBook(2, bookController)
	// if err != nil {
	// 	fmt.Println("Can't Delete This Book")
	// 	return
	// }
}	