package utils

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nsrvel/golang-example/internal/models"
)

type RestErr interface {
	Status() int
	Message() string
	Error() string
}

type RestError struct {
	ErrStatus  int    `json:"status,omitempty"`
	ErrMessage string `json:"message,omitempty"`
	ErrError   error  `json:"-"`
}

func (e RestError) Error() string {
	if e.ErrError != nil {
		return e.ErrError.Error()
	}
	return ""
}

func (e RestError) Status() int {
	return e.ErrStatus
}

func (e RestError) Message() string {
	return e.ErrMessage
}

func ErrorWrapper(statusCode int, message string, err error) RestErr {
	return RestError{
		ErrStatus:  statusCode,
		ErrMessage: message,
		ErrError:   err,
	}
}

func ParseError(err error) RestErr {
	if restErr, ok := err.(RestErr); ok {
		return restErr
	}
	return RestError{
		ErrStatus:  http.StatusInternalServerError,
		ErrMessage: "failed to parse rest error",
	}
}

func WriteErrorResponse(c *fiber.Ctx, statusCode int, message string, err string) error {
	return c.Status(statusCode).JSON(models.Response{
		StatusCode: statusCode,
		Message:    message,
		Errors:     err,
	})
}

func WriteSuccessResponse(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	return c.Status(statusCode).JSON(models.Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	})
}
