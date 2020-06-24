package api

import (
	"github.com/labstack/echo/v4"
)

// handles creating thing upon POST request with json data
func (a *API) createUser(c echo.Context) error {
	var in User
	
	// return json response
	return c.JSON(200, in)
}
