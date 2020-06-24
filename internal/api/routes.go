package api

import "github.com/labstack/echo/v4"

// initiate API routes
func (a *API) initRoutes() {
	r := echo.New()
	// save router to server
	a.srv.Handler = r
	
	routeGroup := r.Group("/api/v1")
	
	// declare routes
	routeGroup.GET("/user", a.getAllUsers)
	routeGroup.POST("/user", a.createUser)
	routeGroup.GET("/user/:id", a.getUser)
	routeGroup.PUT("/user/:id", a.updateUser)
	routeGroup.DELETE("/user/:id", a.deleteUser)
}

