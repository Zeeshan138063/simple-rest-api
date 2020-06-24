package api

import "github.com/labstack/echo/v4"

// initiate API routes
func (a *API) initRoutes() {
	r := echo.New()
	// save router to server
	a.srv.Handler = r
	
	routeGroup := r.Group("/api/v1")
	
	// declare routes
	routeGroup.GET("/user", nil)
	routeGroup.POST("/user", nil)
	routeGroup.GET("/user/:id", nil)
	routeGroup.PUT("/user/:id", nil)
	routeGroup.DELETE("/user/:id", nil)
}

