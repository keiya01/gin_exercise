package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/keiya01/ginsampleapp/src/controllers"
)

func Routes(app *gin.Engine) {
	commonCtrl := new(controllers.CommonController)
	homeCtrl := new(controllers.HomeController)
	userCtrl := new(controllers.UsersController)

	app.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:3000"},
		AllowMethods:  []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:  []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders: []string{"X-Total-Count"},
	}))

	app.GET("/", homeCtrl.IndexCtrl)
	app.POST("/auth", commonCtrl.AuthenticateUser)
	app.POST("/regist", userCtrl.CreateCtrl)
	app.POST("/login", userCtrl.LoginCtrl)
	homeGroup := app.Group("/todos")
	{
		homeGroup.POST("/show", homeCtrl.ShowCtrl)
		homeGroup.POST("/create", homeCtrl.CreateCtrl)
		homeGroup.POST("/update", homeCtrl.UpdateCtrl)
		homeGroup.POST("/delete", homeCtrl.DeleteCtrl)
	}
}
