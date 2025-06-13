package posts

import "github.com/gin-gonic/gin"

type PostsModule struct {
	PostsService
}

func NewPostsModule() *PostsModule {
	return &PostsModule{
		PostsService: NewPostsService(),
	}
}

// RegisterRoutes implements PostsModule.
func (p *PostsModule) RegisterRoutes(r *gin.Engine) {
	panic("unimplemented")
}
