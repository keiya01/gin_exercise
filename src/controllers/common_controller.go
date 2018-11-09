package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/keiya01/ginsampleapp/src/models"
	"log"
	"net/http"
)

type CommonController struct {
}

func (common *CommonController) AuthenticateUser(c *gin.Context) {

	params, err := GetUserParams(c)
	if err != nil {
		log.Println("JSON ERROR: ", err)
		return
	}

	user := new(models.User)
	ok := user.FindUserId(params)

	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"status": "ユーザーが見つかりませんでした。",
			"error":  true,
			"user":   map[string]string{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ログインしています！",
		"error":  false,
		"user":   user,
	})
}

func GetUserParams(c *gin.Context) (post models.UserParams, err error) {

	err = c.BindJSON(&post)
	if err != nil {
		return post, err
	}

	return post, nil
}
