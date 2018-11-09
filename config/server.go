package config

import (
	"github.com/gin-gonic/gin"
)

func Start() error {
	app := setup()
	return app.Run(":8686")
}

func setup() *gin.Engine {
	app := gin.New()
	Routes(app)
	return app
}
