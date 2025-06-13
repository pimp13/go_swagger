package users

import (
	"go_swagger/internal/posts"

	"github.com/gin-gonic/gin"
)

type UserModule struct {
	UsersService
	UsersController
}

func NewUserModule(postsModule *posts.PostsModule) *UserModule {
	// external module (dependency)
	userService := NewUsersService(postsModule)

	// internal module
	userController := NewUsersController(userService)

	return &UserModule{
		UsersService:    userService,
		UsersController: *userController,
	}
}

func (um *UserModule) RegisterRoutes(r *gin.Engine) {
	r.GET("/service", um.UsersController.GetUsersHandler)
}
