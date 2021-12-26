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

func NewSessionRouter(router *gin.RouterGroup) {
	r := SessionRouter{
		db:     datasource.GetDataSource(),
		config: config.GetConfig(),
	}

	router.GET("/sessions", r.GetAllSessions)
	router.GET("/sessions/:id", r.GetSessionById)
	router.POST("/sessions", r.CreateSession)
}

func (r *SessionRouter) GetAllSessions(ctx *gin.Context) {
	sessions := make([]models.SessionModel, 0)
	r.db.Find(&sessions)
	ctx.JSON(http.StatusOK, sessions)
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
		return
	}
	session := &models.SessionModel{
		Session: data,
	}
	s.db.Save(session)
	ctx.JSON(http.StatusOK, session)
}

func (s *SessionRouter) GetSessionById(ctx *gin.Context) {
	id, err := readID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	var session models.SessionModel
	s.db.Find(&session, id)
	ctx.JSON(http.StatusOK, session)
}
