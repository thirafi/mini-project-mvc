package database

import (
	"project_structure/config"
	"project_structure/model"
)

func CreateBook(book *model.Book) error {
	if err := config.DB.Create(book).Error; err != nil {
		return err
	}
	return nil
}

func GetBooks() (books []model.Book, err error) {
	if err = config.DB.Find(&books).Error; err != nil {
		return
	}
	return
}

func GetBook(id uint) (book model.Book, err error) {
	book.ID = id
	if err = config.DB.First(&book).Error; err != nil {
		return
	}
	return
}

func UpdateBook(book *model.Book) error {
	if err := config.DB.Updates(book).Error; err != nil {
		return err
	}
	return nil
}

func DeleteBook(book *model.Book) error {
	if err := config.DB.Delete(book).Error; err != nil {
		return err
	}
	return nil
}
