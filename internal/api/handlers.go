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


// get handler for getting a thing by id upon GET request
func (a *API) getUser(c echo.Context) error {
	var out User
	
	a.db.Where("id = ?", c.Param("id")).Find(&out)
	
	if out.ID == 0 {
		return c.String(http.StatusNotFound, "user not found")
	}
	
	return c.JSON(200, out)
}

// get all handler for getting all things upon GET request
func (a *API) getAllUsers(c echo.Context) error {
	var out []User
	
	a.db.Find(&out)
	
	return c.JSON(200, out)
}
