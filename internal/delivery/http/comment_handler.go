package httphandler

import (
	"net/http"

	"github.com/kodinggo/gb-2-api-story-service/internal/model"
	"github.com/labstack/echo/v4"
)

type CommentHandler struct {
	commentUsecase model.CommentUseCase
}

func InitCommentHandler(commentUseCase model.CommentUseCase) CommentHandler {
	return CommentHandler{commentUsecase: commentUseCase}

}

func (h CommentHandler)RegisterRoute(e *echo.Echo){
	g := e.Group("/comment")
	g.POST("",h.Create)
}

func (h CommentHandler) Create(c echo.Context) error {
	var body model.Comment
	err :=c.Bind(&body)
	if err !=nil{
		return echo.NewHTTPError(http.StatusBadRequest,err.Error())
	}
	comment,err :=  h.commentUsecase.Create(c.Request().Context(),body.User_id,body.Story_id,body.Comment)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError,err.Error())
	}
	response :=Response{
		Data: comment,
	}
	return c.JSON(http.StatusAccepted,response)
}

