package routes

import (
	"github.com/LucasCarioca/oservability/pkg/config"
	"github.com/LucasCarioca/oservability/pkg/datasource"
	"github.com/LucasCarioca/oservability/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"net/http"
)

//SessionRouter router for creating and looking up session information
type SessionRouter struct {
	db     *gorm.DB
	config *viper.Viper
}

func NewSessionRouter(app *gin.Engine) {
	r := SessionRouter{
		db:     datasource.GetDataSource(),
		config: config.GetConfig(),
	}

	app.POST("/api/v1/sessions", r.CreateSession)
}

func (s *SessionRouter) CreateSession(ctx *gin.Context) {
	var data models.Session
	err := ctx.BindJSON(&data)
	if err != nil {
		e := models.Error{
			Message: "please check the data being provided and make sure nothing is missing",
			Code:    "INVALID_SESSION_PAYLOAD",
		}
		ctx.JSON(http.StatusBadRequest, e)
	}
	session := &models.SessionModel{
		Session: data,
	}
	s.db.Save(session)
	ctx.JSON(http.StatusOK, session)
}
