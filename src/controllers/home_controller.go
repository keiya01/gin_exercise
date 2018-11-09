package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/keiya01/ginsampleapp/src/models"
	"log"
	"net/http"
)

type HomeController struct {
}

func (t *HomeController) IndexCtrl(c *gin.Context) {
	todo := new(models.Todo)

	todos := todo.GetAll()
	c.Header("X-Total-Count", "25")
	c.JSON(http.StatusOK, gin.H{
		"title": "やること",
		"todos": todos,
	})
}

func (t *HomeController) ShowCtrl(c *gin.Context) {
	todo := new(models.Todo)

	params, err := GetTodoParams(c)
	if err != nil {
		log.Println("JSON ERROR: ", err)
		return
	}

	showTodo := todo.FindOne(params.ID)
	c.JSON(http.StatusOK, gin.H{
		"title": showTodo.Title,
		"todo":  showTodo,
	})

}

func (t *HomeController) CreateCtrl(c *gin.Context) {
	todo := new(models.Todo)

	params, err := GetTodoParams(c)
	if err != nil {
		log.Println("JSON ERROR: ", err)
		return
	}

	showTodo := todo.TodoCreate(params.Title, params.Text)
	c.JSON(http.StatusOK, gin.H{
		"status": "保存しました！",
		"todo":   showTodo,
	})

}

func (t *HomeController) UpdateCtrl(c *gin.Context) {
	todo := new(models.Todo)

	params, err := GetTodoParams(c)
	if err != nil {
		log.Println("JSON ERROR: ", err)
		return
	}

	updatedTodo := todo.TodoUpdate(params)
	c.JSON(http.StatusOK, gin.H{
		"status": "保存しました！",
		"todo":   updatedTodo,
	})

}

func (t *HomeController) DeleteCtrl(c *gin.Context) {
	todo := new(models.Todo)

	params, err := GetTodoParams(c)
	if err != nil {
		log.Println("JSON ERROR: ", err)
		return
	}

	todo.TodoDelete(params.ID, params.Title)
	c.JSON(http.StatusOK, gin.H{
		"status": "削除しました！",
		"todo":   params,
	})

}

func GetTodoParams(c *gin.Context) (post models.TodoParams, err error) {

	err = c.BindJSON(&post)
	if err != nil {
		return post, err
	}

	return post, nil
}
