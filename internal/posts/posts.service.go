package posts

import (
	"context"
	"fmt"
)

type PostsService interface {
	GetPostByID(ctx context.Context, id uint) error
}

type PostsServiceImpl struct {
	// inject dependency
}

func NewPostsService() PostsService {
	return &PostsServiceImpl{}
}

// GetPostByID implements PostsService.
func (p *PostsServiceImpl) GetPostByID(ctx context.Context, id uint) error {
	fmt.Println("*********** posts service called ***********")
	return nil
}

var _ PostsService = (*PostsServiceImpl)(nil)
