package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/keiya01/ginsampleapp/src/models"
	"log"
	"net/http"
)

type UsersController struct {
}

func (u *UsersController) CreateCtrl(c *gin.Context) {
	user := new(models.User)

	params, err := GetUserParams(c)
	if err != nil {
		log.Println("JSON ERROR: ", err)
		return
	}

	showUser, err := user.UserCreate(params)
	if err != nil {
		log.Println("JSON ERROR: ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "保存しました！",
		"user": map[string]string{
			"UserId": showUser.UserId,
			"Name":   showUser.Name,
		},
	})
}

func (u *UsersController) LoginCtrl(c *gin.Context) {
	user := new(models.User)

	params, err := GetUserParams(c)
	if err != nil {
		log.Println("JSON ERROR: ", err)
		return
	}

	err = user.Login(params)
	if err != nil {
		log.Println("JSON ERROR: ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ログインしました。",
		"user": map[string]string{
			"UserId": user.UserId,
			"Name":   user.Name,
		},
	})

}
