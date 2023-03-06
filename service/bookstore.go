package service

import (
	"bookstore/model"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type BookStoreController interface {
	InsertBook(book model.Book) error
	GetBookByID(ID int, book model.Book) (model.Book, error)
	DeleteBook(ID int, book model.Book) (model.Book, error)
}

type BookStoreDB struct {}

func (b *BookStoreDB) InsertBook(book *[]model.Book, db BookStoreDB) error {
	bookJSON, err := json.Marshal(book)
	if err != nil {
		// panic(err)
		return errors.New("Failed to encode to JSON")
	}
	fmt.Println(bookJSON)

	insertBook := os.WriteFile("book.json", bookJSON, 0644)

	if insertBook != nil {
		return errors.New("Failed write to JSON")
	}
	return nil
}

func (b *BookStoreDB) GetBookByID(ID int, db BookStoreDB) ([]model.Book, error) {
	list_book, err := os.ReadFile("book.json")
	if err != nil {
		return []model.Book{}, errors.New("Failed to read JSON")
	}

	book := []model.Book{}

	err = json.Unmarshal(list_book, &book)
	if err != nil {
		panic(err)
	}

	bookID := []model.Book{}
	for _, v := range book {
		if ID == v.ID {
			bookID = append(bookID, v)
		}
	}
	return bookID, nil
}

func (b *BookStoreDB) DeleteBook(ID int, bs BookStoreDB) error {
	list_book, err := os.ReadFile("book.json")
	if err != nil {
		return errors.New("Failed to read JSON")
	}

	book := []model.Book{}

	err = json.Unmarshal(list_book, &book)
	if err != nil {
		panic(err)
	}

	bookAfterDelete := []model.Book{}

	for _, v := range book {
		if v.ID != ID {
			bookAfterDelete = append(bookAfterDelete, v)
		}
	}

	err = os.Remove("book.json")

	bookJSON, err := json.Marshal(bookAfterDelete)
	if err != nil {
		// panic(err)
		return errors.New("Failed to encode to new JSON")
	}

	insertBook := os.WriteFile("book.json", bookJSON, 0644)
	if insertBook != nil {
		return errors.New("Failed write new JSON")
	}
	return nil
}