package usecase

import (
	"errors"
	"fmt"
	"project_structure/model"
	"project_structure/repository/database"
)

func CreateBook(book *model.Book) error {

	// check title cannot be empty
	if book.Title == "" {
		return errors.New("book title cannot be empty")
	}

	// check creator
	if book.Creator == "" {
		return errors.New("book creator cannot be empty")
	}

	err := database.CreateBook(book)
	if err != nil {
		return err
	}

	return nil
}

func GetBook(id uint) (book model.Book, err error) {
	book, err = database.GetBook(id)
	if err != nil {
		fmt.Println("GetBook: Error getting book from database")
		return
	}
	return
}

func GetListBooks() (books []model.Book, err error) {
	books, err = database.GetBooks()
	if err != nil {
		fmt.Println("GetListBooks: Error getting books from database")
		return
	}
	return
}

func UpdateBook(book *model.Book) (err error) {
	err = database.UpdateBook(book)
	if err != nil {
		fmt.Println("UpdateBook : Error updating book, err: ", err)
		return
	}

	return
}

func DeleteBook(id uint) (err error) {
	book := model.Book{}
	book.ID = id
	err = database.DeleteBook(&book)
	if err != nil {
		fmt.Println("DeleteBook : error deleting book, err: ", err)
		return
	}

	return
}
