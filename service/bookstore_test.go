package service

import (
	"bookstore/model"
	"fmt"
	"testing"
)

func TestBookStoreDB_InsertBook(t *testing.T) {
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

	bookController := BookStoreDB{}
	err := bookController.InsertBook(&new_book, bookController)

	if err != nil {
		fmt.Println("Not as expected ", err)
		return
	}

}