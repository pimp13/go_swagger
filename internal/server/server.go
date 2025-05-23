package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	csrf "github.com/utrack/gin-csrf"
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

	// setup middlewares
	store := cookie.NewStore([]byte("secret-key"))

	r.Use(sessions.Sessions("mySession", store))
	r.Use(csrf.Middleware(csrf.Options{
		Secret: "secret-key",
		ErrorFunc: func(c *gin.Context) {
			c.String(http.StatusNotFound, "CSRF token mismatch")
			c.Abort()
		},
	}))
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"http://127.0.0.1:3000",
		},
		AllowMethods: []string{
			"PUT",
			"PATCH",
			"GET",
			"POST",
			"DELETE",
			http.MethodOptions,
			http.MethodHead,
		},
		AllowHeaders: []string{
			"Origin",
		},
		ExposeHeaders: []string{
			"Content-Length",
			"Content-Type",
		},
		AllowCredentials: true,
	}))

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	s.RegisterRoutes(r)

	if err := r.Run(s.addr); err != nil {
		panic("failed to start server: " + err.Error())
	}
}

// helloHandler test api
//
//	@Summary		helloHandler
//	@Description	test api and sending hello world
//	@Produce		json
//	@Success		200	{string}	string	"hello world"
//	@Router			/hello [GET]
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
//	@Router			/user [GET]
func fetchUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"user": "this is user hello test.",
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
//	@Router			/user/{id} [GET]
func getUserByID(c *gin.Context) {
	id := c.Param("id")
	queryParam := c.Query("enumstring")
	c.JSON(200, gin.H{
		"id":         id,
		"queryParam": queryParam,
		"name":       "John Doe",
	})
}

// @Summary		setCsrfToken
// @Description	setCsrfToken
// @Tags			csrf
// @Accept			json
// @Produce		json
// @Success		200	{string}	string
// @Router			/csrf-token [GET]
func setCsrfToken(c *gin.Context) {
	c.String(200, csrf.GetToken(c))
}

// @Summary		getCsrfToken
// @Description	getCsrfToken
// @Tags			csrf
// @Accept			json
// @Produce		json
// @Success		200	{string}	string
// @Router			/csrf-token [POST]
func getCsrfToken(c *gin.Context) {
	c.String(200, "IS OK!ّ")
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
