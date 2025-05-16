package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	addr string
}

func NewServer(addr string) *Server {
	return &Server{
		addr: addr,
	}
}

func (s *Server) Run() {
	r := gin.Default()

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	s.RegisterRoutes(r)

	if err := r.Run(s.addr); err != nil {
		panic("failed to start server: " + err.Error())
	}
}

// helloHandler نمونه هندلر
//
//	@Summary		helloHandler
//	@Description	test api and sending hello world
//	@Produce		json
//	@Success		200	{string}	string	"hello world"
//	@Router			/hello [get]
func helloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world!",
	})
}

// fetchUser handler
//
//	@Summary		fetchUser
//	@Description	this is handler for fetch users
//	@Produce		json
//	@Tags			users
//	@Success		200	{string}	string	"test message"
//	@Router			/user [get]
func fetchUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"user": "this is user",
	})
}

// getUserByID handle url path with id param
//
//	@Summary		getUserByID
//	@Description	get user by id
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id			path		int		true	"user id"
//	@Param			enumstring	query		string	false	"string enums"	Enums(A, B, C)
//	@Success		200			{object}	UserResponse
//	@Failure		400			{object}	ErrorResponse
//	@Failure		404			{object}	ErrorResponse
//	@Router			/user/{id} [get]
func getUserByID(c *gin.Context) {
	id := c.Param("id")
	queryParam := c.Query("enumstring")
	c.JSON(200, gin.H{
		"id":         id,
		"queryParam": queryParam,
		"name":       "John Doe",
	})
}

// UserResponse response model
type UserResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ErrorResponse error model
type ErrorResponse struct {
	Message string `json:"message"`
}
