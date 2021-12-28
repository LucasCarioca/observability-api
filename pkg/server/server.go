package server

import (
	"fmt"
	"github.com/LucasCarioca/oservability/pkg/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func routesInit(app *gin.Engine) {
	v1Sessions := app.Group("/api/v1/sessions")
	routes.NewSessionRouter(v1Sessions)
}

//Init initializes the service and attaches all routers
func Init(config *viper.Viper) {
	app := gin.Default()
	app.Use(cors.Default())
	routesInit(app)
	host := config.GetString("server.host")
	port := config.GetString("server.port")
	app.Run(fmt.Sprintf("%s:%s", host, port))
}
