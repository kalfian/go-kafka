package services

type CommentServiceContract interface{}

type commentService struct{}

func NewCommentService() CommentServiceContract {
	return &commentService{}
}
