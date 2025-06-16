package middleware

import (
	apperrors "github.com/biryanim/hezzl_tz/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			RespondWithError(c, err.Err)
			c.Abort()
			return
		}
	}
}

func RespondWithError(c *gin.Context, err error) {
	appErr := apperrors.FromError(err)

	switch appErr.Code {
	case apperrors.CodeAlreadyExists:
		c.JSON(http.StatusConflict, appErr)
	case apperrors.CodeNotFound:
		c.JSON(http.StatusNotFound, appErr)
	case apperrors.CodeInvalidInput:
		c.JSON(http.StatusBadRequest, appErr)
	case apperrors.CodeInvalidCreds:
		c.JSON(http.StatusForbidden, appErr)
	default:
		c.JSON(http.StatusInternalServerError, appErr)

	}
}
