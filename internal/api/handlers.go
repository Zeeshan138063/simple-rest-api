package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// handles creating thing upon POST request with json data
func (a *API) createUser(c echo.Context) error {
	var in User
	
	// decode request body json to our object
	err := c.Bind(&in)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	
	err = a.validate.Struct(in)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	
	a.db.Create(&in)
	
	// return json response
	return c.JSON(200, in)
}
