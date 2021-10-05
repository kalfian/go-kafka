package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalfian/go-kafka/app/comment/services"
)

type CommentHandlerContract interface {
	AddComment(c *fiber.Ctx) error
}

type commentHandler struct {
	srv services.CommentServiceContract
}

func NewCommentHandler(srv services.CommentServiceContract) CommentHandlerContract {
	return &commentHandler{
		srv: srv,
	}
}

func (handler *commentHandler) AddComment(c *fiber.Ctx) error {

	return nil
}
