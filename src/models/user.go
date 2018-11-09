package models

import (
	"github.com/jinzhu/gorm"
	"github.com/keiya01/ginsampleapp/config/db"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	gorm.Model
	UserId   string
	Name     string
	Password string
}

type UserParams struct {
	ID       int    `json:"ID"`
	UserId   string `json:"UserId"`
	Name     string `json:"Name"`
	Password string `json:"Password"`
}

func (u User) UserCreate(params UserParams) (User, error) {
	setDB := db.DBInit()
	defer setDB.Close()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)

	if err != nil {
		var user User
		log.Println("=======[ERROR]=======")
		log.Println(err)
		log.Println("=======[END]=======")
		return user, err
	}

	user := User{UserId: params.UserId, Name: params.Name, Password: string(hashedPassword)}
	setDB.Create(&user)

	log.Println("====== Query Results ======")
	log.Println("UserId: ", user.UserId)
	log.Println("Name: ", user.Name)
	log.Println("上記の内容で保存しました。")
	log.Println("====== END ======")

	return user, nil
}

func (u *User) Login(params UserParams) error {
	setDB := db.DBInit()
	defer setDB.Close()

	u.UserId = params.UserId

	setDB.First(&u)

	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(params.Password))
	if err != nil {
		log.Println("=======[ERROR]=======")
		log.Println(params.Password)
		log.Println(err)
		log.Println("=======[END]=======")
		return err
	}

	log.Println("====== Query Results ======")
	log.Println("UserId: ", u.UserId)
	log.Println("Name: ", u.Name)
	log.Println("ログインしました。")
	log.Println("====== END ======")

	return nil
}

func (u User) FindUserId(params UserParams) bool {
	setDB := db.DBInit()
	defer setDB.Close()

	u.UserId = params.UserId
	setDB.First(&u)
	if u.Name == "" {
		log.Println("User not found.")
		return false
	}

	return true
}
