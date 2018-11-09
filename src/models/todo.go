package models

import (
	"github.com/jinzhu/gorm"
	"github.com/keiya01/ginsampleapp/config/db"
	"log"
)

type Todo struct {
	gorm.Model
	Title     string
	Text      string
	Completed bool
}

type TodoParams struct {
	ID        int    `json:"ID"`
	Title     string `json:"Title"`
	Text      string `json:"Text"`
	Completed bool   `json:"Completed"`
}

func (t *Todo) GetAll() []Todo {
	setDB := db.DBInit()
	defer setDB.Close()
	var todos []Todo
	setDB.Order("created_at desc").Find(&todos)

	log.Println("====== Query Results ======")
	log.Println(todos)
	log.Println("====== END ======")

	return todos
}

func (t *Todo) FindOne(id int) Todo {
	setDB := db.DBInit()
	defer setDB.Close()

	var todo Todo
	setDB.First(&todo, id)

	log.Println("====== Query Results ======")
	log.Println(todo)
	log.Println("====== END ======")

	return todo
}

func (t *Todo) TodoCreate(title string, text string) Todo {
	setDB := db.DBInit()
	defer setDB.Close()

	todo := Todo{Title: title, Text: text, Completed: false}
	setDB.Create(&todo)

	log.Println("====== Query Results ======")
	log.Println(todo.Title)
	log.Println(todo.Text)
	log.Println("上記の内容で保存しました。")
	log.Println("====== END ======")

	return todo
}

func (t *Todo) TodoUpdate(params TodoParams) Todo {
	setDB := db.DBInit()
	defer setDB.Close()

	var todo Todo
	setDB.First(&todo, params.ID)
	setDB.Model(&todo).Update("Title", params.Title)
	setDB.Model(&todo).Update("Text", params.Text)
	setDB.Model(&todo).Update("Completed", params.Completed)

	log.Println("====== Query Results ======")
	log.Println("title: ", params.Title)
	log.Println("text: ", params.Text)
	log.Println("completed: ", params.Completed)
	log.Println("上記の内容に変更しました。")
	log.Println("====== END ======")

	return todo

}

func (t *Todo) TodoDelete(id int, title string) {
	setDB := db.DBInit()
	defer setDB.Close()

	setDB.Where("ID = ?", id).Delete(&Todo{})

	log.Println("====== Query Results ======")
	log.Println(title)
	log.Println("上記の内容を削除しました。")
	log.Println("====== END ======")

}
