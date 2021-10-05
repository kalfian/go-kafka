package comment

import (
	"github.com/kalfian/go-kafka/app/comment/handlers"
	"github.com/kalfian/go-kafka/app/comment/services"
)

func ProvideHandler() handlers.CommentHandlerContract {
	services := services.NewCommentService()

	return handlers.NewCommentHandler(services)
}
