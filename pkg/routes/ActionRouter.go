package routes

import (
	"github.com/LucasCarioca/oservability/pkg/config"
	"github.com/LucasCarioca/oservability/pkg/datasource"
	"github.com/LucasCarioca/oservability/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
	"net/http"
)

//ActionRouter router for creating and reading actions
type ActionRouter struct {
	db     *gorm.DB
	config *viper.Viper
}

func NewActionRouter(router *gin.RouterGroup) {
	r := ActionRouter{
		db:     datasource.GetDataSource(),
		config: config.GetConfig(),
	}

	router.GET("/actions", r.GetAllActions)
	router.POST("/actions", r.CreateAction)
}

func (r *ActionRouter) GetAllActions(ctx *gin.Context) {
	actions := make([]models.ActionModel, 0)
	r.db.Find(&actions)
	ctx.JSON(http.StatusOK, actions)
}

func (r *ActionRouter) CreateAction(ctx *gin.Context) {
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
	var c int64
	session := models.SessionModel{}
	r.db.Find(&session, data.SessionId).Count(&c)
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
		Action:  data,
		Session: session,
	}
	r.db.Save(action)
	ctx.JSON(http.StatusOK, action)
}
