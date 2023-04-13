package route

import (
	"net/http"
	"orm_mvc/controllers"
	m "orm_mvc/middleware"
	"os"

	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Routers() *echo.Echo {
	e := echo.New()
	m.Logger(e)
	godotenv.Load()
	dbHost := os.Getenv("SECRET_KEY")
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Please, login here !, localhost:8000/login")
	})
	// Route / to handler function
	e.POST("/users/login", controllers.LoginUserController)
	e.GET("/users", controllers.GetUsersController, echojwt.JWT(([]byte(dbHost))))
	e.GET("/users/:id", controllers.GetUserController, echojwt.JWT(([]byte(dbHost))))
	e.POST("/users", controllers.CreateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController, echojwt.JWT(([]byte(dbHost))))
	e.PUT("/users/:id", controllers.UpdateUserController, echojwt.JWT(([]byte(dbHost))))

	//  book
	e.GET("/books", controllers.GetBooksController, echojwt.JWT(([]byte(dbHost))))
	e.GET("/books/:id", controllers.GetBookController, echojwt.JWT(([]byte(dbHost))))
	e.POST("/books", controllers.CreateBookController, echojwt.JWT(([]byte(dbHost))))
	e.DELETE("/books/:id", controllers.DeleteBookController, echojwt.JWT(([]byte(dbHost))))
	e.PUT("/books/:id", controllers.UpdateBookController, echojwt.JWT(([]byte(dbHost))))

	return e
}
