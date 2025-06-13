package server

import (
	"go_swagger/internal/posts"
	"go_swagger/internal/users"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes(route *gin.Engine) {

	pm := posts.NewPostsModule()

	um := users.NewUserModule(pm)
	um.RegisterRoutes(route)

	// Routes for API v1
	v1 := route.Group("/api")
	{
		v1.GET("/hello", helloHandler)
		v1.GET("/user", fetchUser)
		v1.GET("/user/:id", getUserByID)

		v1.GET("/weather/:city", getWeatherByCity)

		v1.GET("/agify/:name", getAgify)

		v1.GET("/csrf-token", setCsrfToken)
		v1.POST("/csrf-token", getCsrfToken)
	}
}
