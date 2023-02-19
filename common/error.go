package common

import "github.com/gin-gonic/gin"

type BaseError struct {
	Message string
}

func (e *BaseError) Error() string {
	return e.Message
}

func (e *BaseError) Json() *gin.H {
	return &gin.H{
		"message": e.Message,
	}
}

type ValidationError struct {
	BaseError
	Message string
}

type UnauthorizedError struct {
	BaseError
	Message string
}
