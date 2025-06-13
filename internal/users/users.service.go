package users

import (
	"context"
	"fmt"
	"go_swagger/internal/posts"
)

type UsersService interface {
	GetUserByID(ctx context.Context, id uint) error
}

type UsersServiceImpl struct {
	// inject dependency
	*posts.PostsModule
}

func NewUsersService(pm *posts.PostsModule) UsersService {
	return &UsersServiceImpl{
		PostsModule: pm,
	}
}

// GetUserByID implements UsersService.
func (u *UsersServiceImpl) GetUserByID(ctx context.Context, id uint) error {
	fmt.Println("*********** users service called ***********")
	if err := u.PostsModule.PostsService.GetPostByID(ctx, id); err != nil {
		return err
	}
	return nil
}

var _ UsersService = (*UsersServiceImpl)(nil)
