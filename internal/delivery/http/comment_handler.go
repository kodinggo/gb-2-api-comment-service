package httphandler

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/kodinggo/gb-2-api-comment-service/internal/model"
	"github.com/labstack/echo/v4"
)

type CommentHandler struct {
	commentUsecase model.CommentUseCase
}

func InitCommentHandler(commentUsecase model.CommentUseCase) *CommentHandler {
	return &CommentHandler{commentUsecase: commentUsecase}
}

var validate = validator.New()

func (h CommentHandler) RegisterRoute(e *echo.Echo) {
	g := e.Group("/v1/comment")

	g.POST("", h.Create)
	g.PUT("/:id", h.Update)
	g.DELETE("/:id", h.Delete)
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
		Status: "success",
		Data:   comment,
	}
	return c.JSON(http.StatusCreated, response)
}

func (h *CommentHandler) Update(c echo.Context) error {
	var data model.Comment
	err := c.Bind(&data)
	if err != nil {
		return newHTTPError(c, http.StatusBadRequest, err.Error())
	}

	err = validate.Struct(data)
	if err != nil {
		return newHTTPError(c, http.StatusBadRequest, err.Error())
	}

	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	result, err := h.commentUsecase.Update(c.Request().Context(), int64(idInt), &data)
	if err != nil {
		return newHTTPError(c, http.StatusBadRequest, err.Error())
	}

	response := Response{
		Status: "success",
		Data:   result,
	}
	return c.JSON(http.StatusOK, response)
}

func (h *CommentHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = h.commentUsecase.Delete(c.Request().Context(), int64(idInt))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response := Response{
		Status: "success",
	}
	return c.JSON(http.StatusOK, response)
}
