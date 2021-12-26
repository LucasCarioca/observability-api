package routes

import (
	"github.com/LucasCarioca/oservability/pkg/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func readID(ctx *gin.Context) (*int, *models.Error) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return nil, &models.Error{
			Message: "Not a valid id",
			Code:    "INVALID_ID",
		}
	}
	return &id, nil
}
