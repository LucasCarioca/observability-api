package common

import (
	"github.com/LucasCarioca/oservability/pkg/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

//ReadSessionID get the id of the resource
func ReadSessionID(ctx *gin.Context) (*int, *models.Error) {
	id, err := strconv.Atoi(ctx.Param("sessionId"))
	if err != nil {
		return nil, &models.Error{
			Message: "Not a valid id",
			Code:    "INVALID_ID",
		}
	}
	return &id, nil
}

//RespondWithError handles responding to requests with and error message
func RespondWithError(ctx *gin.Context, status int, error models.Error) {
	ctx.JSON(status, error)
	ctx.Abort()
}
