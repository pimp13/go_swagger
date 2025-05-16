package server

import "github.com/gin-gonic/gin"

func (s *Server) RegisterRoutes(route *gin.Engine) {
	// Routes for API v1
	v1 := route.Group("/api")
	{
		v1.GET("/hello", helloHandler)
		v1.GET("/user", fetchUser)
		v1.GET("/user/:id", getUserByID)

		v1.GET("/weather/:city", getWeatherByCity)

		v1.GET("/agify/:name", getAgify)
	}
}
