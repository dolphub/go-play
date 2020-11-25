package apperr

import (
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AppError struct {
	Status  int    `json:"-"`
	Message string `json:"message,omitempty"`
}

func (e AppError) Error() string {
	return e.Message
}

// NewStatus generates new error containing only http status code
func NewStatus(status int) *AppError {
	return &AppError{Status: status}
}

// New generates an application error
func New(status int, msg string) *AppError {
	if msg != "" {

		return &AppError{Status: status, Message: msg}
	}
	return NewStatus(status)
}

func BadRequestError(msg string) *AppError {
	return New(http.StatusBadRequest, msg)
}

func ForbiddenError(msg string) *AppError {
	return New(http.StatusForbidden, msg)
}

func UnauthorizedError(msg string) *AppError {
	return New(http.StatusUnauthorized, msg)
}

func NotFoundError(msg string) *AppError {
	return New(http.StatusNotFound, msg)
}

var validationErrors = map[string]string{
	"required": " is required, but was not received",
	"min":      "'s value or length is less than allowed",
	"max":      "'s value or length is bigger than allowed",
}

func getVldErrorMsg(s string) string {
	if v, ok := validationErrors[s]; ok {
		return v
	}
	return " failed on " + s + " validation"
}

func Response(c *gin.Context, err error) {
	spew.Dump(err)
	switch err.(type) {
	case *AppError:
		e := err.(*AppError)
		if e.Message == "" {
			c.AbortWithStatus(e.Status)
		} else {
			c.AbortWithStatusJSON(e.Status, e)
		}
		return
	case validator.ValidationErrors:
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
}
