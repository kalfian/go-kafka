package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalfian/go-kafka/app/comment/controllers"
)

type CommentHandlerContract interface {
	AddComment(c *fiber.Ctx) error
}

type commentHandler struct {
	c controllers.CommentControllerContract
}

func NewCommentHandler(c controllers.CommentControllerContract) CommentHandlerContract {
	return &commentHandler{
		c: c,
	}
}

func (handler *commentHandler) AddComment(c *fiber.Ctx) error {

	return nil
}
