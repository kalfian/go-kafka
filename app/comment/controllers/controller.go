package controllers

type CommentControllerContract interface{}

type commentController struct{}

func NewCommentController() CommentControllerContract {
	return &commentController{}
}
