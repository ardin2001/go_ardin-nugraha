package response

import "github.com/labstack/echo/v4"

func Response(c echo.Context, statusCode int, message string, data interface{}) error {
	return c.JSON(statusCode, map[string]interface{}{
		"status":   200,
		"error":    "null",
		"messages": message,
		"user":     data,
	})
}
