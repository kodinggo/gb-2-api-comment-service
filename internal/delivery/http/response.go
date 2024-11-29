package httphandler

import "github.com/labstack/echo/v4"

type Response struct {
	Status   string         `json:"status"`
	Message  string         `json:"message"`
	Data     any            `json:"data"`
	Metadata map[string]any `json:"metadata,omitempty"`
}

func newHTTPError(c echo.Context, httpCode int, msg string) error {
	return c.JSON(httpCode, Response{
		Status:  "failed",
		Message: msg,
	})
}
