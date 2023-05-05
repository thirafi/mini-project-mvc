package controller

import (
	"net/http"
	"project_structure/model"
	"strconv"

	"project_structure/usecase"

	"github.com/labstack/echo/v4"
)

func GetBookcontroller(c echo.Context) error {
	books, e := usecase.GetListBooks()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"books":  books,
	})
}

func GetBookController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	book, err := usecase.GetBook(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error get book",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"books":  book,
	})
}

// create new book
func CreateBookController(c echo.Context) error {
	book := model.Book{}
	c.Bind(&book)

	if err := usecase.CreateBook(&book); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error create book",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"book":    book,
	})
}

// delete book by id
func DeleteBookController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := usecase.DeleteBook(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete book",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf buku tidak dapat di hapus",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book",
	})
}

// update book by id
func UpdateBookController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	book := model.Book{}
	c.Bind(&book)
	book.ID = uint(id)
	if err := usecase.UpdateBook(&book); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update book",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf buku tidak dapat di ubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book",
	})
}
