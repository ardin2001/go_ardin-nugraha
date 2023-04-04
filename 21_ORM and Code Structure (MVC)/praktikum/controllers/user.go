package controllers

import (
	"net/http"
	"orm_mvc/config"
	"orm_mvc/models"
	"strconv"

	"github.com/labstack/echo"
)

// get all users
func GetUsersController(c echo.Context) error {
	var users []models.User

	DB, _ := config.InitDB()
	if err := DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	id := c.Param("id")
	user := models.User{}

	DB, _ := config.InitDB()
	if err := DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user id " + id,
		"users":   user,
	})

}

// create new user
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	DB, _ := config.InitDB()
	if err := DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// // delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	id := c.Param("id")
	DB, _ := config.InitDB()

	new_id, _ := strconv.Atoi(id)
	if err := DB.Delete(&models.User{}, new_id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    "data user id " + id + " deleted successfully",
	})
}

// // update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	id := c.Param("id")
	DB, _ := config.InitDB()
	user := models.User{}

	DB.First(&user, id)
	c.Bind(&user)
	if err := DB.Save(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    user,
	})
}
