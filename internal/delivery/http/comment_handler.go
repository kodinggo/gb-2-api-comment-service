package httphandler

import (
	"net/http"
	"strconv"

	"github.com/kodinggo/gb-2-api-story-service/internal/model"
	"github.com/labstack/echo/v4"
)

type CommentHandler struct {
	commentUsecase model.CommentUseCase
}

func InitCommentHandler(commentUseCase model.CommentUseCase) CommentHandler {
	return CommentHandler{commentUsecase: commentUseCase}

}

func (h CommentHandler) RegisterRoute(e *echo.Echo) {
	g := e.Group("/comment")
	g.POST("", h.Create)
	g.PUT("/:id", h.Update)
}

func (h *CommentHandler) Create(c echo.Context) error {
	var body model.Comment
	err := c.Bind(&body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	comment, err := h.commentUsecase.Create(c.Request().Context(), &body)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	response := Response{
		Data: comment,
	}
	return c.JSON(http.StatusAccepted, response)
}

func (h *CommentHandler) Update(c echo.Context) error {
	var data model.Comment
	id := c.Param("id")

	err := c.Bind(&data)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	idInt, _ := strconv.Atoi(id)
	result, err := h.commentUsecase.Update(c.Request().Context(), int64(idInt), &data)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	response := Response{
		Data:    result,
		Message: "Sucessfully updated comment",
	}
	return c.JSON(http.StatusAccepted, response)
}
