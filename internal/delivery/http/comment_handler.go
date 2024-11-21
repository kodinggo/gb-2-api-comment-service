package httphandler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/kodinggo/gb-2-api-comment-service/internal/model"
	"github.com/labstack/echo/v4"
)

type CommentHandler struct {
	commentUsecase model.CommentUseCase
}

func InitCommentHandler(commentUseCase model.CommentUseCase) CommentHandler {
	return CommentHandler{commentUsecase: commentUseCase}

}

var validate *validator.Validate

func (h CommentHandler) RegisterRoute(e *echo.Echo) {
	g := e.Group("/comment")
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
		Data: comment,
	}
	return c.JSON(http.StatusAccepted, response)
}

func (h *CommentHandler) Update(c echo.Context) error {
	var data model.Comment
	id := c.Param("id")
	validate = validator.New()
	err := c.Bind(&data)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = validate.Struct(data)
	if err != nil {
		// Jika validasi gagal, kembalikan pesan error
		errorResponse := make(map[string]string)
		for _, fieldErr := range err.(validator.ValidationErrors) {
			// Tambahkan error per field
			errorResponse[fieldErr.Field()] = fmt.Sprintf("Validation failed on '%s'", fieldErr.Tag())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err)
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

func (h *CommentHandler) Delete(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	err = h.commentUsecase.Delete(c.Request().Context(), int64(idInt))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusAccepted, fmt.Sprintf("sucessfully deleted comment id : %v", idInt))
}
