package comment

import (
	"github.com/kalfian/go-kafka/app/comment/controllers"
	"github.com/kalfian/go-kafka/app/comment/handlers"
)

func ProvideHandler() handlers.CommentHandlerContract {
	controller := controllers.NewCommentController()
	return handlers.NewCommentHandler(controller)
}
