package users

import (
	"log"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	UsersService
}

func NewUsersController(us UsersService) *UsersController {
	return &UsersController{
		UsersService: us,
	}
}

func (uc *UsersController) GetUsersHandler(c *gin.Context) {
	err := uc.UsersService.GetUserByID(c.Request.Context(), 12)
	if err != nil {
		log.Fatal(err)
		return
	}

	c.JSON(200, "is ok")
}
