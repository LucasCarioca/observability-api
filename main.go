package main

import (
	"github.com/LucasCarioca/oservability/pkg/config"
	"github.com/LucasCarioca/oservability/pkg/datasource"
	"github.com/LucasCarioca/oservability/pkg/server"
	"github.com/LucasCarioca/oservability/pkg/utils"
)

func main() {
	env := utils.GetEnv()
	config.Init(env)
	datasource.Init(config.GetConfig())
	server.Init(config.GetConfig())
}
