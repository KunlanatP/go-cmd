package errs

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler() fiber.ErrorHandler {
	return FiberErrorHendler
}

func FiberErrorHendler(c *fiber.Ctx, err error) error {
	if fiberError, ok := err.(*fiber.Error); ok {
		return c.Status(fiberError.Code).JSON(ReplyError{
			ErrorMessage: ErrorMessage{
				Message: fiberError.Message,
			},
		})
	}
	if err, ok := err.(*Error); ok {
		return c.Status(err.Code).JSON(ReplyError{
			ErrorMessage: ErrorMessage{
				Message: err.Message,
			},
		})
	}
	return c.Status(http.StatusInternalServerError).JSON(ReplyError{
		ErrorMessage: ErrorMessage{
			Message: "Internal Server Error",
		},
	})
}

func Invalid(message ...string) error {
	return newError(http.StatusBadRequest, message)
}

func newError(code int, message []string) *Error {
	var msg string
	if len(message) > 0 {
		msg = message[0]
	} else {
		msg = http.StatusText(code)
	}
	return &Error{
		Code:    code,
		Message: msg,
	}
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err *Error) Error() string {
	return err.Message
}

var ErrorNotfoundEmpID = Invalid("Employee not found")
var ErrorIDFormat = Invalid("Wrong id format")
var ErrorEmpIDIsRequired = Invalid("EmpID is required")
