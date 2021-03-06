package routes

import (
	"github.com/LucasCarioca/oservability/pkg/auth"
	"github.com/LucasCarioca/oservability/pkg/common"
	"github.com/LucasCarioca/oservability/pkg/config"
	"github.com/LucasCarioca/oservability/pkg/datasource"
	"github.com/LucasCarioca/oservability/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"net/http"
)

//SessionRouter router for creating and looking up session information
type SessionRouter struct {
	db     *gorm.DB
	config *viper.Viper
}

//NewSessionRouter Creates a new routers for session routes
func NewSessionRouter(router *gin.RouterGroup) {
	r := SessionRouter{
		db:     datasource.GetDataSource(),
		config: config.GetConfig(),
	}

	router.GET("/", auth.CheckAPIKEy, r.GetAllSessions)
	router.POST("/", auth.CheckAPIKEy, r.CreateSession)
	router.GET("/:sessionId", auth.CheckSessionKey, r.GetSessionByID)
	router.GET("/:sessionId/actions", auth.CheckSessionKey, r.GetAllActionsForSession)
	router.POST("/:sessionId/actions", auth.CheckSessionKey, r.CreateAction)
}

//GetAllSessions responds to requests with all sessions
func (r *SessionRouter) GetAllSessions(ctx *gin.Context) {
	sessions := make([]models.SessionModel, 0)
	r.db.Find(&sessions)
	ctx.JSON(http.StatusOK, sessions)
}

//CreateSession Creates a new session and returns that the requesting systemd
func (r *SessionRouter) CreateSession(ctx *gin.Context) {
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
	r.db.Save(session)
	ctx.JSON(http.StatusOK, session)
}

//GetSessionByID Retrieves a specific session by its ID
func (r *SessionRouter) GetSessionByID(ctx *gin.Context) {
	id, idError := common.ReadSessionID(ctx)
	if idError != nil {
		ctx.JSON(http.StatusBadRequest, idError)
		return
	}
	var session models.SessionModel
	r.db.Find(&session, id)
	ctx.JSON(http.StatusOK, session)
}

//GetAllActionsForSession returns all sessions for a provided session
func (r *SessionRouter) GetAllActionsForSession(ctx *gin.Context) {
	id, idError := common.ReadSessionID(ctx)
	if idError != nil {
		ctx.JSON(http.StatusBadRequest, idError)
		return
	}
	actions := make([]models.ActionModel, 0)
	r.db.Preload(clause.Associations).Find(&actions, "session_id = ?", id)
	ctx.JSON(http.StatusOK, actions)
}

//CreateAction creates a new action in the provided session
func (r *SessionRouter) CreateAction(ctx *gin.Context) {
	var data models.Action
	err := ctx.BindJSON(&data)
	if err != nil {
		log.Println(err.Error())
		e := models.Error{
			Message: "please check the data being provided and make sure nothing is missing",
			Code:    "INVALID_ACTION_PAYLOAD",
		}
		ctx.JSON(http.StatusBadRequest, e)
		return
	}
	id, idError := common.ReadSessionID(ctx)
	if idError != nil {
		ctx.JSON(http.StatusBadRequest, idError)
		return
	}
	var c int64
	session := models.SessionModel{}
	r.db.Find(&session, id).Count(&c)
	if c < 1 {
		e := models.Error{
			Message: "actions must contain a valid session id",
			Code:    "SESSION_NOT_FOUND",
		}
		ctx.JSON(http.StatusNotFound, e)
		return
	}
	log.Println(">>>>", session, data)
	action := &models.ActionModel{
		Action:    data,
		SessionID: uint(*id),
		Session:   session,
	}
	r.db.Save(action)
	ctx.JSON(http.StatusOK, action)
}
