package controllers

import (
	"net/http"
	"orm_mvc/config"
	"orm_mvc/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all books
func GetBooksController(c echo.Context) error {
	var books []models.Book

	DB, _ := config.InitDB()
	if err := DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"users":   books,
	})
}

// get book by id
func GetBookController(c echo.Context) error {
	// your solution here
	id := c.Param("id")
	book := models.Book{}

	DB, _ := config.InitDB()
	if err := DB.First(&book, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book id " + id,
		"users":   book,
	})

}

// create new book
func CreateBookController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)

	DB, _ := config.InitDB()
	if err := DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"user":    book,
	})
}

// // delete user by id
func DeleteBookController(c echo.Context) error {
	// your solution here
	id := c.Param("id")
	DB, _ := config.InitDB()

	new_id, _ := strconv.Atoi(id)
	if err := DB.Delete(&models.Book{}, new_id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    "data book id " + id + " deleted successfully",
	})
}

// // // update user by id
func UpdateBookController(c echo.Context) error {
	// your solution here
	id := c.Param("id")
	DB, _ := config.InitDB()
	book := models.Book{}

	DB.First(&book, id)
	c.Bind(&book)
	if err := DB.Save(&book).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    book,
	})
}
