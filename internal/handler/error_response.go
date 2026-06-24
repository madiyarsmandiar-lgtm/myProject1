package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/madiyarsmandiar-lgtm/myProject1/internal/apperrors"
)

func respondWithError(c *gin.Context, err error) {
	var appErr *apperrors.AppError

	if errors.As(err, &appErr) {
		c.JSON(appErr.Code, gin.H{"message": appErr.Message})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
}
