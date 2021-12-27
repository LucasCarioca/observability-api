package auth

import (
	"github.com/LucasCarioca/oservability/pkg/config"
	"github.com/LucasCarioca/oservability/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

//AuthCheck middleware function that validates if the request is authenticated
func CheckAPIKEy(ctx *gin.Context) {
	config := config.GetConfig()
	apiKey := config.GetString("API_KEY")
	requestKey := ctx.Query("api_key")
	if apiKey != requestKey {
		ctx.JSON(http.StatusUnauthorized, models.Error{
			Message: "Request is not authorized",
			Code:    "UNAUTHORIZED",
		})
		ctx.Abort()
	}
}
