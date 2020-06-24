package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// handles creating user upon POST request with json data
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


// get handler for getting a user by id upon GET request
func (a *API) getUser(c echo.Context) error {
	var out User
	
	a.db.Where("id = ?", c.Param("id")).Find(&out)
	
	if out.ID == 0 {
		return c.String(http.StatusNotFound, "user not found")
	}
	
	return c.JSON(200, out)
}

// get all handler for getting all users upon GET request
func (a *API) getAllUsers(c echo.Context) error {
	var out []User
	
	a.db.Find(&out)
	
	return c.JSON(200, out)
}


// update handler for updating a user by id upon PUT request
func (a *API) updateUser(c echo.Context) error {
	var in User
	// decode request body json to our object
	err := c.Bind(&in)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	
	idInt, _ := strconv.Atoi(c.Param("id"))
	in.ID = uint(idInt)
	
	a.db.Model(&in).Updates(in)
	
	var out User
	a.db.Where("id = ?", c.Param("id")).Find(&out)
	
	// return json response
	return c.JSON(200, out)
}

// delete handler for deleting a user by id upon DELETE request
func (a *API) deleteUser(c echo.Context) error {
	a.db.Where("id = ?", c.Param("id")).Delete(&User{})
	
	return c.JSON(200, echo.Map{"message": "user deleted successfully"})
}
