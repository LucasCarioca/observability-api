package auth

import (
	"github.com/LucasCarioca/oservability/pkg/common"
	"github.com/LucasCarioca/oservability/pkg/config"
	"github.com/LucasCarioca/oservability/pkg/datasource"
	"github.com/LucasCarioca/oservability/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

//CheckAPIKEy middleware function that validates if the request is authenticated
func CheckAPIKEy(ctx *gin.Context) {
	config := config.GetConfig()
	apiKey := config.GetString("API_KEY")
	requestKey := ctx.Query("api_key")
	if apiKey != requestKey {
		common.RespondWithError(ctx, http.StatusUnauthorized, models.Error{
			Message: "Request is not authorized",
			Code:    "UNAUTHORIZED",
		})
	}
}

//CheckSessionKey middleware function that validates if the request is for the session provided
func CheckSessionKey(ctx *gin.Context) {
	id, idError := common.ReadSessionID(ctx)
	if idError != nil {
		common.RespondWithError(ctx, http.StatusBadRequest, *idError)
		ctx.Abort()
	}
	var session models.SessionModel
	datasource.GetDataSource().Find(&session, id)
	requestKey := ctx.Query("api_key")
	if requestKey != session.SessionKey {
		common.RespondWithError(ctx, http.StatusUnauthorized, models.Error{
			Message: "Request is not authorized",
			Code:    "UNAUTHORIZED",
		})
	}
}
